# [Go Class: 01 Hello world!](https://www.youtube.com/watch?v=A9HfEhvpOEY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=2)

## Online Go Playground

- [Go Playground](https://play.golang.org/)
- [Repl.it](https://repl.it/)

## Hello World!

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
- In go, every program has to have a main function. It tells go where to start.
- It also has to have a package main. It tells go that this is an executable program.

## Running the program
To run a program anywhere, type `go run <filename>.go`

To run a program in the current directory, you can simply type `go run .`

When running the program, Go actually compiles the program and then runs it.
