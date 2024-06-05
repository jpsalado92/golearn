package main

import "fmt"

func main() {
    v := make([]int, 5) // len(v)=5, cap(v)=5
    fmt.Printf("len=%d cap=%d %v\n", len(v), cap(v), v)
    fmt.Printf("Memory address for v: %p\n", &v)
    fmt.Printf("Memory address for v[0]: %p\n", &v[0])
    
    v = append(v, 0) // len(v)=6, cap(v)=10
    fmt.Printf("len=%d cap=%d %v\n", len(v), cap(v), v)
    fmt.Printf("Memory address for v: %p\n", &v)
    fmt.Printf("Memory address for v[0]: %p\n", &v[0])
}