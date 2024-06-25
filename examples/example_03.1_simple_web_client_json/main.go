package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// In this example we are using a simple web client that reads a json response
// from a server and parses it into a struct

const url = "https://jsonplaceholder.typicode.com"

type todo struct {
	// userId	int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close() // Close so that the socket can be reused

	if resp.StatusCode == http.StatusOK { // If 200
		body, err := io.ReadAll(resp.Body) // Read the whole body of the response

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		// Unmarshal the json response into the struct
		my_todo := todo{}
		err = json.Unmarshal(body, &my_todo)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		fmt.Printf("%#v\n", my_todo)

	}
}
