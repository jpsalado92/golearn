package main

import (
	"fmt"
	"my_project"
    "os"
)

func main() {
    names := os.Args[1:]
    fmt.Printf(my_project.Say(names))
}