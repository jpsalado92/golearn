# Go 1.22.3 Quick Reference

- [Go 1.22.3 Quick Reference](#go-1223-quick-reference)
	- [Links to resources](#links-to-resources)
	- [Reference](#reference)
		- [Go main characteristics](#go-main-characteristics)
		- [Operators and punctuation](#operators-and-punctuation)
		- [Keywords](#keywords)
		- [Predeclared identifiers](#predeclared-identifiers)
	- [Practical reference](#practical-reference)
		- [Showing the type and value of something](#showing-the-type-and-value-of-something)
		- [Simplest program](#simplest-program)
		- [Testing](#testing)

## Links to resources

- [Go Documentation - Standard library](https://pkg.go.dev/std)
- [Go Playground](https://play.golang.org/)
- [Go Keywords](https://golang.org/ref/spec#Keywords)
- [Go Operators and punctuation](https://golang.org/ref/spec#Operators_and_punctuation)
- [Go Predeclared identifiers](https://golang.org/ref/spec#Predeclared_identifiers)
- [Repl.it](https://repl.it/)

## Reference

### Go main characteristics

- It is a compiled language, not an interpreted language.
- It is a statically typed language, which means that the type of a variable is known at compile time.
- It is a garbage-collected language, which means that the memory is managed automatically.
- It is a concurrent language, which means that it is easy to write programs that do many things at once.
- It is a simple language, which means that it is easy to learn and easy to read.

### Operators and punctuation
```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

### Keywords
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Predeclared identifiers
``` 
Types:
	any bool byte comparable
	complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap clear close complex copy delete imag len
	make max min new panic print println real recover
```

## Practical reference
### Showing the type and value of something

```go
fmt.Printf("%T %[1]v\n", a)

```

### Simplest program

- In go, every program has to have a main function. It tells go where to start.
- It also has to have a package main. It tells go that this is an executable program.
- To run a program anywhere, type `go run <filename>.go` or `go run .` in the current directory.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

### Testing
In order to everything to work properly, you must create a root folder for your project, in this case we call it `my_project`.

The folder structure should be like this:
```
my_project
├── cmd
│   └── main.go
├── go.mod
├── hello.go
└── hello_test.go
```

cmd/main.go
```go
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
```

hello.go
```go
package my_project

import (
	"fmt"
	"strings"
)

func Say(names []string)string {
	if len(names) == 0 {
		return "Hello, world!"
	}
	return fmt.Sprintf("Hello, %s.", strings.Join(names, ", ") + "!")
}
```

hello_test.go
```go
package my_project

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, test!."
	got := Say([]string{"test"})

	if got != want {
		t.Errorf("SayHello() = %q, wanted %q", got, want)
	}
}

func TestMultipleSayHello(t *testing.T) {
	subtests := []struct {
		names []string
		want  string
	}{
		{[]string{"test"}, "Hello, test!."},
		{names: []string{"test1", "test2"}, want: "Hello, test1, test2!."},
		{[]string{"test1", "test2", "test3"}, "Hello, test1, test2, test3!."},
		{[]string{}, "Hello, world!"},
	} 

	for _, tt := range subtests {
		t.Run(tt.want, func(t *testing.T) {
			got := Say(tt.names)
			if got != tt.want {
				t.Errorf("SayHello() = %q, wanted %q", got, tt.want)
			}
		})
	}
}

```

For the file `go.mod` (you must run `go mod init my_project` in the root folder)

In order to run tests, you can simply run `go test` in the root folder.