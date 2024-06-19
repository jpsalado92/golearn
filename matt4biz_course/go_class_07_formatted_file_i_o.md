# [Go Class: 07 Formatted & File I/O](https://www.youtube.com/watch?v=dqEtGT-dxoY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=9)
- [Go Class: 07 Formatted \& File I/O](#go-class-07-formatted--file-io)
  - [Standard Input, Output and Error](#standard-input-output-and-error)
  - [Print functions in go](#print-functions-in-go)
    - [Printing to Stdout](#printing-to-stdout)
    - [Printing to to anything with a Write method](#printing-to-to-anything-with-a-write-method)
    - [Printing to a string with Sprint](#printing-to-a-string-with-sprint)
  - [Formatting](#formatting)
    - [Reusing arguments](#reusing-arguments)
    - [Printing numbers](#printing-numbers)
    - [Different types and values](#different-types-and-values)
  - [File I/O](#file-io)
    - [Building `cat` in go with os.Stdout](#building-cat-in-go-with-osstdout)
    - [Building `cat` in go with ioutil.ReadAll](#building-cat-in-go-with-ioutilreadall)
    - [Building `wc` (word counter) in go](#building-wc-word-counter-in-go)

## Standard Input, Output and Error

Reading from Stdin
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Println("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello", name)
}
```

Printing to Stdout
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Printing to Stderr
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprintln(os.Stderr, "Error occurred")
}
```


## Print functions in go

### Printing to Stdout


```go
package main

import "fmt"

func main() {
    fmt.Print("Hello, ")
    fmt.Print("World!\n")
    fmt.Println("Hello, World!")
    fmt.Printf("Hello, %s!\n", "World")
}
```

### Printing to to anything with a Write method

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprint(os.Stdout, "Hello, ")
    fmt.Fprint(os.Stdout, "World!\n")
    fmt.Fprintln(os.Stdout, "Hello, World!")
    fmt.Fprintf(os.Stdout, "Hello, %s!\n", "World")
}
```

### Printing to a string with Sprint

```go
package main

import (
    "fmt"
)

func main() {
    s := fmt.Sprint("Hello, ", "World!\n")
    s = fmt.Sprintln("Hello, World!")
    s = fmt.Sprintf("Hello, %s!\n", "World")
}
```

## Formatting

Resource: [List of all format codes](https://pkg.go.dev/fmt)

### Reusing arguments

```go
package main

import (
    "fmt"
)

func main() {
	var r []rune
    // Stating same argument twice
	fmt.Printf("value: `%v` type: `%T`\n", r, r) // value: `[]` type: `[]int32`
    // Reusing arguments
	fmt.Printf("value: `%v` type: `%[1]T`\n", r) // value: `[]` type: `[]int32`
}
```

### Printing numbers

```go
package main

import (
	"fmt"
)

func main() {
	a, b := 10, 20.1234

	// Print the values of a and b with different precisions
	fmt.Printf("a = %d, b = %.0f\n", a, b)
	fmt.Printf("a = %d, b = %.1f\n", a, b)
	fmt.Printf("a = %d, b = %.2f\n", a, b)
	fmt.Printf("a = %d, b = %.3f\n", a, b)
	fmt.Printf("a = %d, b = %f\n", a, b)

	// Print the values of a and b with spacings
	fmt.Printf("|%6d|%6f\n", a, b)  // 6 spaces
	fmt.Printf("|%06d|%06f\n", a, b)  // 6 spaces with leading zeros
	fmt.Printf("|%-6d|%-6f\n", a, b)  // 6 spaces with left alignment
}
```

### Different types and values

```go
package main

import (
	"fmt"
)

func main() {

	// A string
	s := "Hello, 世界"
	fmt.Printf("%T\n", s)  // string
	fmt.Printf("%q\n", s)  // Hello, 世界
	fmt.Printf("%v\n", s)  // Hello, 世界
	fmt.Printf("%#v\n", s)  // "Hello, 世界"

	// A byte-string (a slice of bytes)
	b := []byte(s)
	fmt.Printf("%T\n", b)  // []uint8
	fmt.Printf("%q\n", b)  // "Hello, 世界"
	fmt.Printf("%v\n", b)  // [72 101 108 108 111 44 32 228 184 150 231 149 140]
	fmt.Printf("%v\n", string(b))  // Hello, 世界

	// Slice of integers
	i := []int{1, 2, 3}
	fmt.Printf("%T\n", i)  // []int
	fmt.Printf("%v\n", i)  // [1 2 3]
	fmt.Printf("%#v\n", i)  // []int{1, 2, 3}

	// Array of runes
	a := [3]rune{'a', 'b', 'c'}
	fmt.Printf("%T\n", a)  // [3]int32
	fmt.Printf("%q\n", a)  // ['a' 'b' 'c']
	fmt.Printf("%v\n", a)  // [97 98 99]
	fmt.Printf("%#v\n", a)  // [3]int32{97, 98, 99}
	
	// Map of strings to integers
	m := map[string]int{"one": 1, "two": 2}
	fmt.Printf("%T\n", m)  // map[string]int
	fmt.Printf("%v\n", m)  // map[one:1 two:2]
	fmt.Printf("%#v\n", m)  // map[string]int{"one":1, "two":2}
}
```



## File I/O

- Package `os` provides functions to open, read, write and close files. Also provides functions to create, remove, rename and list files.
- Package `io` provides interfaces to read and write data.
- Package `io/ioutil` provides utility functions to read/write entire files to to/from memory.
- Package `bufio` provides buffered I/O.
- Package `strconv` provides functions to convert from string representations.

### Building `cat` in go with os.Stdout

```go

package main

import (
    "io"
    "fmt"
    "os"
)

func main() {
    for _, filename := range os.Args[1:] {
        file, err := os.Open(filename)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }
        if _, err:=io.Copy(os.Stdout, file); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        file.Close()
    }
}
```

### Building `cat` in go with ioutil.ReadAll

```go

package main

import (
    "fmt"
    "os"
    "io"
)

func main() {
    for _, filename := range os.Args[1:] {
        file, err := os.Open(filename)

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }
        data, err := io.ReadAll(file)
        if ; err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }

        file.Close()
    }
}
```

### Building `wc` (word counter) in go

```go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	total_wc := 0
    for _, filename := range os.Args[1:] {
        var line_count, word_count, characer_count int

        file , err := os.Open(filename)

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }

        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            line_count++
			total_wc++
			line := scanner.Text()
            word_count += len(line)
            characer_count += len(strings.Fields(line))
        }

		print("line_count: ", line_count, "\n")
		print("word_count: ", word_count, "\n")
		print("characer_count: ", characer_count, "\n")
    }
	print("total_wc: ", total_wc, "\n")
}
```
