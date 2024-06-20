package main

import "fmt"

type address struct {
	city  string
	state string
}

type person struct {
	name    string
	age     int
	address *address
}

func main() {
	address := address{
		city:  "New York",
		state: "NY",
	}
	p := person{
		name:    "Alice",
		age:     25,
		address: &address,
	}
	fmt.Println(p)
	fmt.Println(p.address.city)
	fmt.Println(p.address.state)
}
