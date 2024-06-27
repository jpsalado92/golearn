package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

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

func main() {

	// Load from xkcd.json
	var comics []ComicDescription
	file, err := os.Open("../xkcd.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&comics)
	if err != nil {
		panic(err)
	}

	var matchedComics []ComicDescription

	// Unpack search terms into an array
	searchTerms := os.Args[1:]

comic_loop:
	for _, comic := range comics {
		// If any of the searchTerms are part of the transcript, add to matchedComics
		for _, searchTerm := range searchTerms {
			if strings.Contains(comic.Transcript, searchTerm) {
				matchedComics = append(matchedComics, comic)
				continue comic_loop
			}
		}
	}

	// Order matchedComics by number
	sort.Slice(matchedComics, func(i, j int) bool {
		return matchedComics[i].Num < matchedComics[j].Num
	})

	fmt.Printf("SearchTerms: %v\n", searchTerms)
	fmt.Printf("Matched Comics:\n")
	for _, comic := range matchedComics {
		fmt.Printf("    %d: %50s\t[ https://xkcd.com/%[1]d ]\n", comic.Num, comic.Title)
	}
}
