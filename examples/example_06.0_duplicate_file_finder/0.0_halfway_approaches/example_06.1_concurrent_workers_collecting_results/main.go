package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func scanDirPaths(dirPath string) ([]pathMetadata, error) {
	pathsPayload := make([]pathMetadata, 0)

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

		pathsPayload = append(pathsPayload, pathMetadata{
			path:       path,
			name:       d.Name(),
			size:       info.Size(),
			modifiedAt: info.ModTime(),
			isDir:      d.IsDir(),
			isRegular:  d.Type().IsRegular(),
		})
		return nil
	})

	return pathsPayload, nil
}

type pathMetadata struct {
	path       string
	name       string
	size       int64
	modifiedAt time.Time
	isDir      bool
	isRegular  bool
	md5        []byte
}

func printPathMetadata(pm pathMetadata) {
	fmt.Printf("\033[33mPath: %s\033[0m\n", pm.path)
	fmt.Printf(" - Name:        %s\n", pm.name)
	fmt.Printf(" - Size:        %d bytes\n", pm.size)
	fmt.Printf(" - ModifiedAt:  %s\n", pm.modifiedAt)
	fmt.Printf(" - IsDir:       %t\n", pm.isDir)
	fmt.Printf(" - IsRegular:   %t\n", pm.isRegular)
	md5String := hex.EncodeToString(pm.md5)              // Convert to hex string
	fmt.Printf(" - md5:         %s\n\n", md5String[:32]) // Display first 10 characters of hex string
}

func calculateMD5(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory path")
		os.Exit(1)
	}

	dirPath := os.Args[1]
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Printf("Directory `%s` does not exist\n", dirPath)
		os.Exit(1)
	}

	fmt.Printf("Scanning directory `%s`\n", dirPath)
	pathsMetadata, err := scanDirPaths(os.Args[1])
	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d files\n", len(pathsMetadata))

	processFiles(pathsMetadata)
}

func processFiles(pathsMetadata []pathMetadata) {
	j := 0

	for i := range pathsMetadata {
		if pathsMetadata[i].isDir || !pathsMetadata[i].isRegular {
			continue
		}
		md5Hash, err := calculateMD5(pathsMetadata[i].path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", pathsMetadata[i].path, err)
			continue
		}
		pathsMetadata[i].md5 = md5Hash
		printPathMetadata(pathsMetadata[i])
		j++
	}

	fmt.Printf("Processed %d files\n", j)
}
