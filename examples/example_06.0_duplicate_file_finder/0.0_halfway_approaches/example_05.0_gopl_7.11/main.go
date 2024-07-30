package main

import (
	"io"
	"log"
	"net/http"
)

type item struct {
	name  string
	price string
}

var items []item

func handleItemCreate(w http.ResponseWriter, req *http.Request) {
	item := item{
		name:  req.FormValue("name"),
		price: req.FormValue("price"),
	}
	for _, i := range items {
		if i.name == item.name {
			http.Error(w, "Item already exists", http.StatusConflict)
			return
		}
	}
	items = append(items, item)
}
func handleItemDelete(w http.ResponseWriter, req *http.Request) {
	for i, item := range items {
		if item.name == req.FormValue("name") {
			items = append(items[:i], items[i+1:]...)
			return

		}
	}
	http.NotFound(w, req)
}

func handleItemRead(w http.ResponseWriter, req *http.Request) {
	for _, item := range items {
		if item.name == req.FormValue("name") {
			io.WriteString(w, item.name+" "+item.price+"\n")
			return
		}
	}
	http.NotFound(w, req)
}
func handleRead(w http.ResponseWriter, req *http.Request) {
	for _, item := range items {
		io.WriteString(w, item.name+" "+item.price+"\n")
	}
}

func handleItemUpdate(w http.ResponseWriter, req *http.Request) {
	for i, item := range items {
		if item.name == req.FormValue("name") {
			items[i].price = req.FormValue("price")
			return
		}
	}
	http.NotFound(w, req)
}

func main() {
	http.HandleFunc("/createItem", handleItemCreate)
	http.HandleFunc("/read", handleRead)
	http.HandleFunc("/readItem", handleItemRead)
	http.HandleFunc("/updateItem", handleItemUpdate)
	http.HandleFunc("/deleteItem", handleItemDelete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
