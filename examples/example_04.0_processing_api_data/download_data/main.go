package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

var base_url = "https://xkcd.com"
var comic_description_endpoint = "/info.0.json"

type ComicDescription struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safe_title string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type ComicResult struct {
	Comic ComicDescription
	Error error
}

var numberOfComics = 2500

func main() {
	comicChannel := make(chan ComicResult, numberOfComics)
	var wg sync.WaitGroup


	for i := 1; i < numberOfComics; i++ {
		wg.Add(1)
		go func(comicNumber int) {
			defer wg.Done()
			comic, shouldReturn := getComic(comicNumber)
			if shouldReturn {
				comicChannel <- ComicResult{Error: fmt.Errorf("error fetching comic %d", comicNumber)}
				return
			}
			comicChannel <- ComicResult{Comic: comic}
		}(i)
	}

	go func() {
		wg.Wait()
		close(comicChannel)
	}()

	var comicSlice []ComicDescription
	for result := range comicChannel {
		if result.Error == nil {
			comicSlice = append(comicSlice, result.Comic)
		} else {
			fmt.Println(result.Error)
		}
	}

	fileName := fmt.Sprintf("%d_comics.json", numberOfComics)
	filePath := "../" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode the JSON data and write it to the file
	jsonData, err := json.Marshal(comicSlice)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file successfully")
}

func getComic(i int) (ComicDescription, bool) {
	target_url := base_url + "/" + fmt.Sprint(i) + comic_description_endpoint
	content, err := http.Get(target_url)
	fmt.Println(target_url)
	if err != nil {
		fmt.Println(err)
		return ComicDescription{}, true
	}

	defer content.Body.Close()

	var comic ComicDescription
	err = json.NewDecoder(content.Body).Decode(&comic)
	if err != nil {
		fmt.Println(err)
		return ComicDescription{}, true
	}
	return comic, false
}

