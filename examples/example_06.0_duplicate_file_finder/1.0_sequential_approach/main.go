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
	"time"
)

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

type filePathMetadataType struct {
	path       string
	name       string
	size       int64
	modifiedAt time.Time
}

func getFilePaths(directoryPath string) map[string][]filePathMetadataType {
	fileHashMap := make(map[string][]filePathMetadataType)
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
		fileMeta := filePathMetadataType{
			path:       path,
			name:       d.Name(),
			size:       info.Size(),
			modifiedAt: info.ModTime(),
		}
		filePathHash := calculateFilePathHash(fileMeta.path)
		fileHashMap[filePathHash] = append(fileHashMap[filePathHash], fileMeta)
		return nil
	}

	err := filepath.WalkDir(
		directoryPath,
		visitFunc,
	)

	if err != nil {
		log.Panic(err)
	}

	return fileHashMap
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

func printFilePathMetadata(pm filePathMetadataType) {
	fmt.Printf("\033[33mPath: %s\033[0m\n", pm.path)
	fmt.Printf(" - Name:        %s\n", pm.name)
	fmt.Printf(" - Size:        %d bytes\n", pm.size)
	fmt.Printf(" - ModifiedAt:  %s\n", pm.modifiedAt)
	fmt.Println()
}

func main() {

	directoryPath := getAssertDirectoryPath()
	fileHashMap := getFilePaths(directoryPath)

	for _, files := range fileHashMap {
		if len(files) > 1 {
			fmt.Printf("\033[31mFound %d files with the same hash\033[0m\n", len(files))
			for _, file := range files {
				printFilePathMetadata(file)
			}
		}
	}
}
