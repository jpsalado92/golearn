# [Go Class: 02 Simple Example](https://www.youtube.com/watch?v=-EYNVEv-snE&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=3)
- [Go Class: 02 Simple Example](#go-class-02-simple-example)
	- [Example 1: Program that reads from STDIN and writes to STDOUT](#example-1-program-that-reads-from-stdin-and-writes-to-stdout)
	- [Example 2: Program that reads args and writes to STDOUT](#example-2-program-that-reads-args-and-writes-to-stdout)
	- [Example 3: Program that uses a function in another file and a test for it](#example-3-program-that-uses-a-function-in-another-file-and-a-test-for-it)

## Example 1: Program that reads from STDIN and writes to STDOUT

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println("You entered:", text)
}
```

## Example 2: Program that reads args and writes to STDOUT

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("You entered:", os.Args[1])
}
```

## Example 3: Program that uses a function in another file and a test for it

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