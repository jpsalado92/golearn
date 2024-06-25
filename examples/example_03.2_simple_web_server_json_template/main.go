package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	// "io"
)

var form = `
<h1>Todo #{{.Id}}</h1>
<div>{{printf "User %d" .UserId}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>
`

type todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"

	resp, err := http.Get(base + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// body, err := io.ReadAll(resp.Body)

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 	return
		// }

		item := todo{}

		if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		tmpl, err := template.New("todo").Parse(form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		err = tmpl.Execute(w, item)

		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

	}
}

func main() {
	http.HandleFunc("/", handler)                // Bind the handler to the top-leve route
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080
}
