package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Encoding to JSON and storing in a file
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}

	fmt.Println("Encoded JSON:", string(jsonData))

	// Store the JSON data in a file
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file successfully")

	// Read the JSON data in a file
	file, err = os.Open("person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	jsonData = make([]byte, 100)
	n, err := file.Read(jsonData)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Read data from file:", string(jsonData[:n]))

	// Decoding JSON data

	var newPerson Person
	err = json.Unmarshal(jsonData[:n], &newPerson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded JSON:", newPerson)
}
