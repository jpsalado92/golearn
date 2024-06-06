package main

import "fmt"

func main() {
    a := []int{1,2,3}
    b := a[:1]
    c := b[:2]
    d := c[0:1:1]
    
    a[0] = 10
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("a: %v", a), cap(a), len(a))
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("b: %v", b), cap(b), len(b))
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("c: %v", c), cap(c), len(c))
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("d: %v", d), cap(d), len(d))
}