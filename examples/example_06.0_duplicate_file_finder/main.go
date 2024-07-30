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

func getFilesToCompare(dirPath string) ([]filePathMetadataType, error) {
	err := validateDirectory(os.Args[1])
	if err != nil {
		fmt.Printf("Error validating directory: %v\n", err)
		return nil, err
	}
	fmt.Printf("Scanning directory `%s`\n", dirPath)

	filePathsPayload := make([]filePathMetadataType, 0)

	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		info, err := d.Info()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() || !d.Type().IsRegular() {
			return nil
		}
		if info.Size() == 0 {
			return nil
		}

		filePathsPayload = append(filePathsPayload, filePathMetadataType{
			path:       path,
			name:       d.Name(),
			size:       info.Size(),
			modifiedAt: info.ModTime(),
		})
		return nil
	})
	fmt.Printf("Found %d files\n", len(filePathsPayload))
	return filePathsPayload, nil
}

type filePathMetadataType struct {
	path       string
	name       string
	size       int64
	modifiedAt time.Time
}
type md5Result struct {
	filePathMetadata filePathMetadataType
	md5              string
}

func printFilePathMetadata(pm filePathMetadataType) {
	fmt.Printf("\033[33mPath: %s\033[0m\n", pm.path)
	fmt.Printf(" - Name:        %s\n", pm.name)
	fmt.Printf(" - Size:        %d bytes\n", pm.size)
	fmt.Printf(" - ModifiedAt:  %s\n", pm.modifiedAt)
}

func calculateMD5(filePath string) string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func validateDirectory(dirPath string) error {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory `%s` does not exist", dirPath)
	}
	if err != nil {
		return fmt.Errorf("error accessing directory `%s`: %v", dirPath, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("path `%s` is not a directory", dirPath)
	}
	return nil
}

type md5FileMap map[string][]filePathMetadataType

func processFiles(pathsMetadata []filePathMetadataType) md5FileMap {
	md5Reference := make(md5FileMap)
	ch := make(chan md5Result)

	for i := range pathsMetadata {
		go calculateMD5Hash(ch, pathsMetadata[i])
	}

	for range pathsMetadata {
		res := <-ch
		if v, ok := md5Reference[res.md5]; !ok {
			md5FileSlice := make([]filePathMetadataType, 0)
			md5FileSlice = append(md5FileSlice, res.filePathMetadata)
			md5Reference[res.md5] = md5FileSlice
		} else {
			md5Reference[res.md5] = append(v, res.filePathMetadata)
		}
	}
	return md5Reference
}

func calculateMD5Hash(ch chan<- md5Result, filePathMetadata filePathMetadataType) {
	md5Hash := calculateMD5(filePathMetadata.path)
	ch <- md5Result{filePathMetadata, md5Hash}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory path")
		os.Exit(1)
	}

	filePathMetadata, err := getFilesToCompare(os.Args[1])
	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		os.Exit(1)
	}

	myMap := processFiles(filePathMetadata)
	for k, v := range myMap {
		// if error pass
		if k == "" {
			continue
		}
		if len(v) > 1 {
			fmt.Printf("\033[31mMD5: %s\033[0m\n", k)
			for _, file := range v {
				printFilePathMetadata(file)
				println()
			}
		}
	}
}
