# [Go Class: 35 Benchmarking](https://www.youtube.com/watch?v=nk4rALKLQkc)

- [Go Class: 35 Benchmarking](#go-class-35-benchmarking)
  - [Benchmarking in Go](#benchmarking-in-go)
  - [Defining a benchmark function](#defining-a-benchmark-function)
    - [Define what to actually measure with b.ResetTimer()](#define-what-to-actually-measure-with-bresettimer)
  - [Collecting benchmark results](#collecting-benchmark-results)
    - [Time limit](#time-limit)
    - [CPU cores](#cpu-cores)
    - [Memory allocation](#memory-allocation)

## Benchmarking in Go

The same way we can use the `go test` command to run tests, we can use the `go test`
command to run benchmarks.

Some basic rules around benchmarks in Go are:

- Benchmarks live in test files ending with `_test.go`
- Benchmarks are run with `go test -bench=. ./..`
- Functions to be benchmarked must start with `Benchmark`

Benchmarking in Go offers many features, including:

- Benchmarking performance
- Benchmarking memory allocation
- Benchmarking based on number of CPU cores

## Defining a benchmark function

To define a benchmark function, we need to create a function that starts with `Benchmark`
followed by the name of the function we want to benchmark. For example, if we want to
benchmark a function called `Fibonacci`, we can create a benchmark function called
`BenchmarkFibonacci`.

```go
package main

import (
    "testing"
)

func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(10)
    }
}
```

### Define what to actually measure with b.ResetTimer()

When benchmarking a function, we want to measure the time it takes to run the function
and not the time it takes to set up the benchmark. To do this, we can use the `b.ResetTimer()`
function to reset the timer before running the function we want to benchmark.

```go
func BenchmarkFibonacci(b *testing.B) {
    // Set up variables
    n := 10
    // Reset the timer
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Fibonacci(n)
    }
}
```

Like this we are not measuring the time it takes to set up the variables, but only the time
it takes to run the `Fibonacci` function.

## Collecting benchmark results

The benchmkark results are collected in a table format that looks like this:

```bash
goos: darwin
goarch: amd64
BenchmarkFibonacci-8   	1000000000	         0.000000 ns/op
PASS
ok  	command-line-arguments	0.000s
```

The table shows the following columns:

- `goos`: The operating system.
- `goarch`: The architecture.
- `BenchmarkFibonacci-8`: The name of the benchmark, including the **number of CPU** cores used.
- `1000000000`: The number of iterations.
- `0.000000 ns/op`: The time taken per iteration.

### Time limit

With no set time limit, the benchmark will run for a default of 1 second. We can set a
time limit by using the `-benchtime` flag. For example, to run the benchmark for 10
seconds, we can use the following command:

```bash
go test -bench=. -benchtime=10s
```

### CPU cores

By default, Go will run benchmarks using only one CPU core. To run benchmarks using
multiple CPU cores, we can use the `-cpu` flag. For example, to run benchmarks using 4 CPU
cores, we can use the following command:

```bash
go test -bench=. -cpu=4
```

### Memory allocation

To benchmark memory allocation, we can use the `-benchmem` flag:

```bash
go test -bench=. -benchmem
```

This will show the memory allocated per operation:

```bash
goos: darwin
goarch: amd64
BenchmarkFibonacci-8   	1000000000	         0.000000 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	command-line-arguments	0.000s
```
