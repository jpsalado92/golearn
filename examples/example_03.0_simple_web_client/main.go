package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// In this example we are using a simple web client that sends a request to a server

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <name>")
		os.Exit(-1)
	}
	resp, err := http.Get("http://localhost:8080/" + os.Args[1])

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
		fmt.Println(string(body))
	}
}
