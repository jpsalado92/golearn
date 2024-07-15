# [Go Class: 25 Context](https://www.youtube.com/watch?v=0x_oUlxzw5A&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=26)

- [Go Class: 25 Context](#go-class-25-context)
	- [Contexts](#contexts)
	- [Context values](#context-values)
	- [Working with contexts](#working-with-contexts)
		- [Defining a context](#defining-a-context)
		- [Using a context with a request](#using-a-context-with-a-request)
	- [Conventions](#conventions)
		- [`ctx` as first argument of a function](#ctx-as-first-argument-of-a-function)
		- [Defer `cancel`](#defer-cancel)
		- [Use `select` when dealing with channels](#use-select-when-dealing-with-channels)
		- [Avoid using optional values in a context](#avoid-using-optional-values-in-a-context)
	- [Examples](#examples)
		- [Example 1: Parallel GET with timeout context](#example-1-parallel-get-with-timeout-context)
		- [Example 2.1: Parallel GET, get first response (bad)](#example-21-parallel-get-get-first-response-bad)
		- [Example 2.2: Parallel GET, get first response (good)](#example-22-parallel-get-get-first-response-good)
		- [Example 3: Middleware with private context value key](#example-3-middleware-with-private-context-value-key)
		- [Example 4: Log with trace ID](#example-4-log-with-trace-id)

## Contexts

Package context defines the `Context` type, which carries deadlines, cancellation signals,
and other request-scoped values across API boundaries and between processes.

Many network or database requests, for example, take a context for cancellation.

A context offers **two controls:**

- A channel that closes when the cancellation occurs.
- An error that is readable once the channel closes.

A context is a tree of `immutable` nodes which can be extended.
Cancellation or timeout applies to the current context and its subrtree.
A subtree may be created with a shorter timeout (but not longer)

## Context values

A context may also carry request-specific values, such as a trace ID or authentication
related data.

Package specific types should define their own context keys, or use a type to avoid collisions.

```go
type key string

const myKey key = "myKey"

func WithMyKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, myKey, value)
}

func GetMyKey(ctx context.Context) string {
	if value, ok := ctx.Value(myKey).(string); ok {
		return value
	}
	return ""
}
```


## Working with contexts

### Defining a context

```go

// Start with an empty context
ctx := context.Background()

// Add a value to the context
ctx = context.WithValue(ctx, "key", "value")

// Access the value in the context
value := ctx.Value("key")

// Add a timeout or a deadline to the context
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
```

In the example above, the first `ctx := context.Background()` creates the parent context,
from there on, more layers are added to that context.

The difference between `WithTimeout` and `WithDeadline` is that `WithTimeout` takes a
duration, while `WithDeadline` takes a set time, like `May 1st, 2022 at 5:00 PM`.

### Using a context with a request

```go
// Get a request working with the context
req, err := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
resp, err := http.DefaultClient.Do(req)
```

The `http.NewRequestWithContext` function is used to create a request with a context.
Then, the `http.DefaultClient.Do(req)` function is used to make the request.

```go

```

## Conventions

### `ctx` as first argument of a function

Contexts should be passed as the first argument of a function as `ctx context.Context`.

```go
func DoSomething(ctx context.Context) {
    // Do something
}
```

### Defer `cancel`

The `cancel` function should be deferred to ensure that the context is always cancelled.
Not doing so, may cause a memory leak.

```go

func DoSomething(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// Do something
}
```

One can verify if the context is cancelled by using a `select` block.

```go
// Check if the context is cancelled
select {
case <-ctx.Done():
    fmt.Println("Cancelled")
default:
    fmt.Println("Not cancelled")
}
```

### Use `select` when dealing with channels

When dealing with channels, use a `select` block to check if the context is cancelled.
This is necessary so that if a parent timeout is set, the child context is also cancelled.

```go
select {
	case r := <-results:
			return &r, nil
	case <-ctx.Done():
		fmt.Println("Cancelled")
}
```

### Avoid using optional values in a context

Avoid using optional values in a context, as it may lead to confusion.

## Examples

### Example 1: Parallel GET with timeout context

In this example, a timeout was injected to each GET request.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"log"
	"runtime"
)

func get(
	ctx context.Context,
	url string,
	ch chan<- result,
) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	if resp, err := http.DefaultClient.Do(req); err != nil {
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
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	results := make(chan result)
	for _, url := range urls {
		go get(ctx, url, results)
	}

	for range urls {
		res := <-results
		if res.err != nil {
			fmt.Printf("%-20s %s\n", res.url, res.err)
		} else {
			fmt.Printf("%-20s %s\n", res.url, res.latency)
		}
	}
}

```

### Example 2.1: Parallel GET, get first response (bad)

This one is useful for cases in which you have multiple available redundant microservices,
and you want to get the first response of any.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"log"
	"runtime"
)

func get(
	ctx context.Context,
	url string,
	ch chan<- result,
) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	if resp, err := http.DefaultClient.Do(req); err != nil {
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

func first(ctx context.Context, urls []string) (*result, error) {
	results := make(chan result)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, url := range urls {
		go get(ctx, url, results)
	}

	select {
	case r := <-results:
		return &r, nil
	case <-ctx.Done():
		// This should be included in case the parent context had a timeout.
		return nil, ctx.Err()
	}
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
	}

	r, _ := first(context.Background(), urls)

	if r.err != nil {
		fmt.Printf("%-20s %s\n", r.url, r.err)
	} else {
		fmt.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(9 * time.Second)
	log.Println("quit anyway...", runtime.NumGoroutine(), "still running")
}

```

### Example 2.2: Parallel GET, get first response (good)

The example above does not take care of the many goroutines that are still running after the first response is received.
This example takes care of that by using a buffered channel.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"log"
	"runtime"
)

func get(
	ctx context.Context,
	url string,
	ch chan<- result,
) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	if resp, err := http.DefaultClient.Do(req); err != nil {
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

func first(ctx context.Context, urls []string) (*result, error) {
	results := make(chan result, len(urls))
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, url := range urls {
		go get(ctx, url, results)
	}

	select {
	case r := <-results:
		return &r, nil
	case <-ctx.Done():
		// This should be included in case the parent context had a timeout.
		return nil, ctx.Err()
	}
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
	}

	r, _ := first(context.Background(), urls)

	if r.err != nil {
		fmt.Printf("%-20s %s\n", r.url, r.err)
	} else {
		fmt.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(9 * time.Second)
	log.Println("quit anyway...", runtime.NumGoroutine(), "still running")
}

```


### Example 3: Middleware with private context value key

In this example, a middleware is created to add a traceId to a request handler.

```go
package main

import (
	"context"
	"net/http"
)

type contextKey string

const TraceKey contextKey = 1

func AddTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if traceId := r.Header.Get("X-Trace-Id"); traceId != "" {
			ctx = context.WithValue(ctx, TraceKey, traceId)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
```

### Example 4: Log with trace ID

In this example, a log is created with a trace ID.

```go
package main

type contextKey string

const TraceKey contextKey = 1

func ContextLogger(ctx context.Context, f string, args ...interface{}) {

	traceId, ok := ctx.Value(TraceKey).(string)
	if ok && traceId != "" {
		f = fmt.Sprintf("[%s] %s", traceId, f)
	}
	log.Printf(f, args...)
}
