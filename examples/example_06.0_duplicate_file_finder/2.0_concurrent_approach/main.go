package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type results map[string][]filePathMetadataType

type filePathMetadataType struct {
	path       string
	name       string
	size       int64
	modifiedAt time.Time
}

func printFilePathMetadata(pm filePathMetadataType) {
	fmt.Printf("\033[33mPath: %s\033[0m\n", pm.path)
	fmt.Printf(" - Name:        %s\n", pm.name)
	fmt.Printf(" - Size:        %d bytes\n", pm.size)
	fmt.Printf(" - ModifiedAt:  %s\n", pm.modifiedAt)
	fmt.Println()
}

type pair struct {
	hash         string
	filePathMeta filePathMetadataType
}

func getAssertDirectoryPath() string {

	// Check if directory was provided as argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <directory_path>")
		os.Exit(1)
	}

	// Check if directory exists
	_, err := os.Stat(os.Args[1])
	if os.IsNotExist(err) {
		fmt.Println("Directory does not exist")
		os.Exit(1)
	}
	return os.Args[1]
}
func calculateFilePathHash(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return ""
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Printf("Error calculating hash: %v\n", err)
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))

}
func searchTree(directoryPath string, filePathMetadataChan chan<- filePathMetadataType) error {
	visitFunc := func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			fmt.Println(err)
			return err
		}

		if d.IsDir() {
			return nil
		} // Skip directories

		info, err := d.Info()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.Size() < 1024*1024/2 {
			return nil
		} // skip files smaller than 0,5MB
		filePathMetadataChan <- filePathMetadataType{
			path:       path,
			name:       d.Name(),
			size:       info.Size(),
			modifiedAt: info.ModTime(),
		}
		return nil
	}
	return filepath.WalkDir(directoryPath, visitFunc)
}

func getPairs(filePathMetaDataChan <-chan filePathMetadataType, pairChan chan<- pair, doneChan chan<- bool) {
	for fileMeta := range filePathMetaDataChan {
		hash := calculateFilePathHash(fileMeta.path)
		pairChan <- pair{
			hash:         hash,
			filePathMeta: fileMeta,
		}
	}
	doneChan <- true
}

func collectHashMap(pairs <-chan pair, resultChan chan<- results) {
	fileHashMap := make(results)

	for pair := range pairs {
		fileHashMap[pair.hash] = append(fileHashMap[pair.hash], pair.filePathMeta)
	}

	resultChan <- fileHashMap
}
func run(dir string) results {
	// Determine the number of worker goroutines to use, based on the number of CPU cores
	workers := 2 * runtime.GOMAXPROCS(0)

	// Create channels for file metadata, pairs, completion signals, and results
	filePathMetaDataChan := make(chan filePathMetadataType)
	pairChan := make(chan pair)
	doneChan := make(chan bool)
	resultChan := make(chan results)

	// Start worker goroutines to process file metadata and generate pairs
	for i := 0; i < workers; i++ {
		go getPairs(filePathMetaDataChan, pairChan, doneChan)
	}

	// Start a goroutine to collect the results into a hash map
	go collectHashMap(pairChan, resultChan)

	// Traverse the directory tree and send file metadata to the filePathMetaDataChan
	if err := searchTree(dir, filePathMetaDataChan); err != nil {
		log.Fatal(err)
	}

	// Close the filePathMetaDataChan to signal that no more file metadata will be sent
	close(filePathMetaDataChan)

	// Wait for all worker `getPairs` goroutines to finish processing
	for i := 0; i < workers; i++ {
		<-doneChan
	}

	// Close the pairChan to signal that no more pairs will be sent
	close(pairChan)

	// Return the final results collected from the resultChan
	return <-resultChan
}

func main() {
	directoryPath := getAssertDirectoryPath()

	if hashes := run(directoryPath); hashes != nil {
		for _, files := range hashes {
			if len(files) > 1 {
				fmt.Printf("\033[31mFound %d files with the same hash\033[0m\n", len(files))
				for _, file := range files {
					printFilePathMetadata(file)
				}
			}
		}
	}
}
