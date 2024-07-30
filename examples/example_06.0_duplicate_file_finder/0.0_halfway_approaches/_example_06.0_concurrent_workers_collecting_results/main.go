package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"
)

type pathPayload struct {
	path string
	name string
}

func createPathPayload() []pathPayload {
	pls := make([]pathPayload, 0)

	filepath.WalkDir("../../", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if d.IsDir() {
			return nil
		} // skip directories

		info, err := d.Info()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.Size() < 1024*1024/2 {
			return nil
		} // skip files smaller than 0,5MB


		pls = append(pls, pathPayload{
			path: path,
			name: d.Name(),
		})
		return nil
	})

	return pls
}

func pathPayloadProcessor(p pathPayload) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Processing task at `%s`\n", p.path)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task at `%s` done\n", p.path)
}

func main() {

	// payloadChan := make(chan payload, 4)

	paths := createPathPayload()
	// fmt.Println(pls)
	for i := range paths {
		// fmt.Println(pls[i])
		go pathPayloadProcessor(paths[i])
	}
	time.Sleep(5 * time.Second)
}
