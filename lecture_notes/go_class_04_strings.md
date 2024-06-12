# [Go Class: 04 Strings](https://www.youtube.com/watch?v=nxWqANttAdA&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=5)

## Strings and runes

* Strings in Go are all unicode.
* Strings in Go are immutable. You can't change a string once it's created.
* Raw strings use backticks, like `string with "quotes"`, and they don't evaluate escape sequences like `\n` or `\t`.

* A `rune` is the Go equivalent of a character.
* `rune` is a type which is an alias for `int32`, and it takes up 4 bytes of memory. This makes it possible to represent all unicode characters.
* Runes are enclosed in single quotes, like `'a'`.

## UTF-8 encoding

* When you store a string in Go, it's stored as a sequence of bytes.
* Go uses UTF-8 encoding to store unicode strings.
* UTF-8 is a variable-length encoding scheme that uses 1 to 4 bytes to represent a character.

## Common strings methods

* `len(s string) int` returns the length of the string in bytes.
* `len([]rune(s string)) int` returns the length of the string in runes.
* `len([]byte(s string)) int` returns the length of the string in bytes.
* `string(r rune) string` converts a rune to a string.
* `[]rune(s string) []rune` converts a string to a slice of runes.
* `[]byte(s string) []byte` converts a string to a slice of bytes.
* `s[i]` returns the byte at index `i` in the string.
* `s[i:j]` returns the substring starting at index `i` and ending at index `j-1`.

## Examples

```go	
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

```

## String manipulation

* You can't add elements to a string, as strings are immutable.
* You can use the `+` operator to concatenate strings.
* You can use the `fmt.Sprintf` function to format strings.
* You can use the `strings` package to manipulate strings.
* You can use the `bytes` package to manipulate strings.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Concatenation
	s1 := "Hello"
	s2 := "World"
	s3 := s1 + " " + s2
	fmt.Println(s3) // Hello World
	// Formatting
	s4 := fmt.Sprintf("%s %s", s1, s2)
	fmt.Println(s4) // Hello World
	// Manipulation
	s5 := "Hello, World"
	s6 := strings.Replace(s5, "World", "世界", -1)
	fmt.Println(s6) // Hello, 世界
}
```
