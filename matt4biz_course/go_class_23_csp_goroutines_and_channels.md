# [Go Class: 23 CSP, Goroutines, and Channels](https://www.youtube.com/watch?v=zJd7Dvg3XCk&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=23)

- [Go Class: 23 CSP, Goroutines, and Channels](#go-class-23-csp-goroutines-and-channels)
  - [What is a Go routine?](#what-is-a-go-routine)
  - [What are channels in Go?](#what-are-channels-in-go)
  - [Examples](#examples)
    - [Example 1: Sum of a slice](#example-1-sum-of-a-slice)
    - [Example 2: Parallel GET](#example-2-parallel-get)
    - [Example 3.1: Stream of IDs with possible race condition](#example-31-stream-of-ids-with-possible-race-condition)
    - [Example 3.2: Stream of IDs with no race condition](#example-32-stream-of-ids-with-no-race-condition)
    - [Example 4: Prime Sieve](#example-4-prime-sieve)

## What is a Go routine?

It is a unit of independent execution (function) that runs concurrently with other functions.

A Go routine is not a thread. The number of threads used by the Go runtime is limited, and the Go runtime manages the scheduling of goroutines on these threads.

In order to create a goroutine, one must use the `go` keyword followed by the function call.

```go
go someFunction()
```

One must make sure that Go routines are not blocked, as this can cause memory leaks.

## What are channels in Go?

A channel is a one-way communication mechanism that allows goroutines to communicate with each other.

It is a typed conduit through which you can send and receive values with the channel operator, `<-`.

A write operation will always happen before a read operation, and the read operation will block until a value is available to read. So the channel can be used to synchronize goroutines.

Channels are first-class values, just like strings or integers, and can be passed as arguments to functions or returned from functions.

One can think of channels as the pipe operator in Linux, where the output of one command is the input of another command.

Go allows multiple reads and writes to a channel, but it is important to note that the reads and writes are blocking operations.

## Examples

### Example 1: Sum of a slice

```go

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)  // Create a channel
    go sum(s[:len(s)/2], c)  // Sum the first half of the slice
    go sum(s[len(s)/2:], c) // Sum the second half of the slice
    x, y := <-c, <-c // x will be the sum of the first half, y the sum of the second half

    fmt.Println(x, y, x+y)
}
```

In the above example, the main function will block on the statement `x, y := <-c, <-c` until the channel `c` has received two values.
After running the two goroutines, the main function will receive the sum of the first half of the slice in `x` and the sum of the second half of the slice in `y`.
Then, it will print the sum of the two halves of the slice.

### Example 2: Parallel GET

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

type result struct {
	url     string
	err     error
	latency time.Duration
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
	}
	results := make(chan result)
	for _, url := range urls {
		go get(url, results) // Start a CSP (Communicating Sequential Process)
	}

	for range urls { // Read from the channel
		res := <-results
		if res.err != nil {
			fmt.Printf("%-20s %s\n", res.url, res.err)
		} else {
			fmt.Printf("%-20s %s\n", res.url, res.latency)
		}
	}
}
```

### Example 3.1: Stream of IDs with possible race condition

```go

package main

var nextID = 0

func handler(w  http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ID: %d\n", nextID)
    // Potential race condition between a READ and a WRITE
    // Another request could be reading the same nextID while this request is writing to it
    nextID++
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Example 3.2: Stream of IDs with no race condition

```go

package main

import (
    "fmt"
    "net/http"
    "sync"
)

var nextID = make(chan int)

func handler(w  http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ID: %d\n", <-nextID)
}

func idGenerator() {
    id := 0
    for {
        nextID <- id
        id++
    }
}

func main() {
    go idGenerator()
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Example 4: Prime Sieve

In this example we will implement a prime number sieve using goroutines and channels.

```go
package main

import "fmt"

func generate(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int, prime int) {
	for i := range in { // Read from the channel until it is closed
		if i%prime != 0 {
			out <- i
		}
	}
	close(out)
}

func sieve(limit int) {
	ch := make(chan int) // Create a channel
	go generate(limit, ch)      // Start the generator
	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		fmt.Println(prime)
	}
}

func main() {
	sieve(100)
}
```
