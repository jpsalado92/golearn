package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database struct {
	data map[string]dollars
	lock sync.Mutex
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	price_str := req.URL.Query().Get("price")

	if _, ok := db.data[name]; ok {
		msg := fmt.Sprintf("Item `%s` already exists", name)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	price, err := strconv.ParseFloat(price_str, 64)
	if err != nil {
		msg := fmt.Sprintf("Invalid price `%s`", price_str)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	db.data[name] = dollars(price)

	fmt.Fprintf(w, "Item `%s` created with price `%s`\n", name, db.data[name])
}
func (db *database) read(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if _, ok := db.data[name]; !ok {
		http.Error(w, "Item does not exist", http.StatusNotFound)
		return
	}
	// io.WriteString(w, name+": "+db[name].String()+"\n")
	fmt.Fprintf(w, "%s: %s\n", name, db.data[name])

}
func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.data {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if _, ok := db.data[name]; !ok {
		http.Error(w, "Item does not exist", http.StatusNotFound)
		return
	}
	delete(db.data, name)
	fmt.Fprintf(w, "Item `%s` deleted\n", name)
}
func (db *database) update(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	price_str := req.URL.Query().Get("price")
	if _, ok := db.data[name]; !ok {
		msg := fmt.Sprintf("Item `%s` does not exist", name)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	price, err := strconv.ParseFloat(price_str, 64)
	if err != nil {
		msg := fmt.Sprintf("Invalid price `%s`", price_str)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	db.data[name] = dollars(price)
	fmt.Fprintf(w, "Item `%s` updated with price `%s`\n", name, db.data[name])
}

func main() {
	db := database{
		data: map[string]dollars{
			"shoes": 50,
			"socks": 5,
			"pants": 30,
			"shirt": 20,
			"hat":   15,
		},
		lock: sync.Mutex{},
	}
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
