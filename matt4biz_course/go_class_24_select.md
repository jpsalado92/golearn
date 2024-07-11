# [Go Class: 24 Select](https://www.youtube.com/watch?v=tG7gII0Ax0Q&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=25)

- [Go Class: 24 Select](#go-class-24-select)
  - [Select](#select)
  - [Example 1: Reading two channels](#example-1-reading-two-channels)
    - [Bad implementation](#bad-implementation)
    - [Good implementation](#good-implementation)
  - [Example 2: Implementing a timeout](#example-2-implementing-a-timeout)
  - [Example 3: Select with periodic timer](#example-3-select-with-periodic-timer)
  - [Example 4: Select with default](#example-4-select-with-default)

## Select

- Select is a control structure in Go that allows you to wait on multiple channel operations (multiplex).
- It can be both a channel that can be read from or written to.
- A select blocks until one of its cases can run, then it executes that case.
- If a default case is present, it will execute that case if no other case is ready. It is advisable to avois using default on loops.
- Most of the times selects run in a loop.
- We can define timeouts by using time.After().

## Example 1: Reading two channels

### Bad implementation

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	chans := []chan int{make(chan int), make(chan int)}
	for i := range chans {
		go func(i int, ch chan int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}
	for i := 0; i < 12; i++ {
		fmt.Println("received", <-chans[0])
		fmt.Println("received", <-chans[1])
	}
}
```

### Good implementation

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	chans := []chan int{make(chan int), make(chan int)}
	for i := range chans {
		go func(i int, ch chan int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}
	for i := 0; i < 12; i++ {
		select {
		case m0 := <-chans[0]:
			fmt.Println("received", m0)
		case m1 := <-chans[1]:
			fmt.Println("received", m1)
		}
	}
}
```

## Example 2: Implementing a timeout

```go
package main

import (
	"log"
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
	stopper := time.After(time.Millisecond * 100)
	list := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com"}
	results := make(chan result)

	for _, url := range list {
		go get(url, results) // start a CSP process
	}
	for range list { // read from the channel
		select {
		case r := <-results:
			log.Printf("%-20s %s\n", r.url, r.latency)
		case <-stopper:
			log.Fatal("timeout")
		}
	}
}
```

## Example 3: Select with periodic timer

```go
package main

import (
	"log"
	"time"
)

const tickRate = 2 * time.Second

func main() {
	log.Println("start")
	ticker := time.NewTicker(tickRate).C // periodic
	stopper := time.After(5 * tickRate)  // one shot
loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}
	log.Println("finish")
}
```

## Example 4: Select with default

```go
package main

import (
  "log"
  "time"
)

func main() {
  ch := make(chan int)
  go func() {
    time.Sleep(2 * time.Second)
    ch <- 1
  }()
  select {
  case <-ch:
    log.Println("received")
  default:
    log.Println("default")
  }
}
```
