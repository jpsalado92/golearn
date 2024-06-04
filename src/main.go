package main

import "fmt"


func main() {
    s:= make([]func(), 4) // create a slice of 4 functions

    for i := 0; i < 4; i++ {
        s[i] = func() {
            fmt.Printf("%d @ %p\n", i, &i) // print the value of i
        }
    }
    for i := 0; i < 4; i++ {
        s[i]() // call each function
    }
}