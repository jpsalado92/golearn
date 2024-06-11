package main

import (
	"fmt"
)

func main() {
	// Rune
	r := 'a'
	fmt.Printf("Type of r: `%T` Value of s: `%[1]v` \n", r)
	fmt.Printf("Length of r: %d \n", len(string(r)))
	// String
	s := "Hello, 世界"
	fmt.Printf("Type of s: `%T` Value of s: `%[1]v` \n", s)
	fmt.Printf("Length of s: %d \n", len(s)) // 13, as there are 13 bytes
	// Slice of runes
	sr := []rune(s)
	fmt.Printf("Type of sr: `%T` Value of sr: `%[1]v` \n", sr)
	fmt.Printf("Length of sr: %d \n", len(sr)) // 9, as there are 9 runes
	// Slice of bytes (compressed)
	sb := []byte(s)
	fmt.Printf("Type of sb: `%T` Value of sb: `%[1]v` \n", sb)
	fmt.Printf("Length of sb: %d \n", len(sb)) // 13, as there are 13 bytes
}
