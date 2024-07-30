package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) create(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	price_str := req.URL.Query().Get("price")

	if _, ok := db[name]; ok {
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
	db[name] = dollars(price)

	fmt.Fprintf(w, "Item `%s` created with price `%s`\n", name, db[name])
}
func (db database) read(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if _, ok := db[name]; !ok {
		http.Error(w, "Item does not exist", http.StatusNotFound)
		return
	}
	// io.WriteString(w, name+": "+db[name].String()+"\n")
	fmt.Fprintf(w, "%s: %s\n", name, db[name])

}
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if _, ok := db[name]; !ok {
		http.Error(w, "Item does not exist", http.StatusNotFound)
		return
	}
	delete(db, name)
	fmt.Fprintf(w, "Item `%s` deleted\n", name)
}
func (db database) update(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	price_str := req.URL.Query().Get("price")
	if _, ok := db[name]; !ok {
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
	db[name] = dollars(price)
	fmt.Fprintf(w, "Item `%s` updated with price `%s`\n", name, db[name])
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
