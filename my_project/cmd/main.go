package main

import (
	"fmt"
)

func main() {
	// Rune
	r := 'a'
	fmt.Printf("Type of r: `%T` Value of s: `%[1]v` \n", r)
	// String
	s := "Hello, 世界"
	fmt.Printf("Type of s: `%T` Value of s: `%[1]v`\n", s)
	// Slice of runes
	sr := []rune(s)
	fmt.Printf("Type of sr: `%T` Value of sr: `%[1]v` \n", sr)
    // Slice of bytes (compressed)
    sb := []byte(s)
    fmt.Printf("Type of sb: `%T` Value of sb: `%[1]v` \n", sb)

}
