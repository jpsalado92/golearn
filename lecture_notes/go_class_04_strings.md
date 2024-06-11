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

## Playground

```go	
package main

import (

    "fmt"

)

func main() {

    // Strings
    s := "Hello, 世界"
    fmt.Println(s)
    fmt.Printf("%T\n", r)

    // Runes
    r := 'a'
    fmt.Println(r)
    fmt.Printf("%T\n", r)

}
```
