
package main

import (
	"fmt"
	"net/http"
	"log"
)

// In this example we are using a concurrent web server that can deal with multiple
// requests at the same time. Run the server and open a browser to
// http://localhost:8080/ to see the output.

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! from %s\n", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)  // Bind the handler to the top-leve route
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080
}