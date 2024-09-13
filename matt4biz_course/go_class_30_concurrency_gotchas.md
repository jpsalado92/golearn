# [Go Class: 30 Concurrency Gotchas](https://www.youtube.com/watch?v=K1hwpNnCJgY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=31)

- [Go Class: 30 Concurrency Gotchas](#go-class-30-concurrency-gotchas)
  - [Gotcha 1: Race conditions](#gotcha-1-race-conditions)
  - [Gotcha 2: Dead locks](#gotcha-2-dead-locks)
    - [Example 2.1: Hanging channel](#example-21-hanging-channel)
    - [Example 2.2: Unlocked locks](#example-22-unlocked-locks)
    - [Example 2.3: The hungry philosophers problem](#example-23-the-hungry-philosophers-problem)
  - [Gotcha 3: Goroutine leaks](#gotcha-3-goroutine-leaks)
  - [Gotcha 4: Incorrect use of WaitGroup](#gotcha-4-incorrect-use-of-waitgroup)
  - [Gotcha 5: Closure capture](#gotcha-5-closure-capture)
  - [Gotcha 6: Select can lead to mistakes](#gotcha-6-select-can-lead-to-mistakes)
    - [Example 6.1: Skipping a channel due to a default case](#example-61-skipping-a-channel-due-to-a-default-case)
    - [Example 6.2: Selecting a nil channel](#example-62-selecting-a-nil-channel)
    - [Example 6.3: Sending to a channel that is full is skipped](#example-63-sending-to-a-channel-that-is-full-is-skipped)
    - [Example 6.4: Beware of using done channels to exit goroutines](#example-64-beware-of-using-done-channels-to-exit-goroutines)
    - [Available channels are selected at random!](#available-channels-are-selected-at-random)
  - [Four considerations when using concurrency](#four-considerations-when-using-concurrency)

## Gotcha 1: Race conditions

A race condition occurs when the outcome of a program depends on the timing or interleaving of multiple threads or processes accessing a shared resource.

In simpler terms, it's like multiple people trying to edit a document at the same time without coordinating. The final version might be a mishmash of changes, leading to unexpected or incorrect results. This can happen in software when multiple threads or processes are trying to modify the same data simultaneously, and the outcome depends on the order in which they do so.

In the following example, requests directed to `/` lead to the handler function printing the value of `nextID`, doing some processing that takes time, and incrementing it.

The race condition occurs when multiple requests are made simultaneously, and the value of `nextID` is read and incremented by multiple goroutines at the same time. This can lead to unexpected results, such as two goroutines reading the same value of `nextID` and incrementing it, resulting in the same value being returned for both requests.

```go
var nextID = 0

func handler(w http.ResponseWriter, r *http.Request) {
    id := nextID
    fmt.Fprintf(w, "<h1>You got %v<h1>", id)
    // Something that takes some time...
    nextID = id + 1  // Race condition
}

func main() {
    http.HandleFunc("/", handler)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
```

## Gotcha 2: Dead locks

A deadlock in concurrent programming occurs when two or more threads (or goroutines, in the context of Go) are unable to proceed with their execution because each is waiting for the other to release a resource or complete an action. This creates a cycle of dependencies that prevents any of the involved threads from making progress.

**Example Scenario**

1. Thread A locks Resource 1 and waits for Resource 2.
1. Thread B locks Resource 2 and waits for Resource 1.
1. Both threads are now waiting indefinitely for each other to release the resources, resulting in a deadlock.

### Example 2.1: Hanging channel

In this example, a deadlock occurs because the main goroutine is waiting to receive from a channel that no other goroutine is sending to:

```go
func main() {
    ch := make(chan bool)
    go func(ok bool) {
        fmt.Println("STARTED")
        if ok {
            ch <- ok
        }
        }(false)
    <-ch  // main goroutine is blocked here
    fmt.Println("DONE")
}
```

### Example 2.2: Unlocked locks

Another possible source of errors in go is related to locking a mutex and then failing to unlock it afterwards. In order to avoid this, it is recommended to use defer to unlock the mutex.

```go
var m sync.Mutex
done := make(chan bool)
go func() {
    m.Lock()
    // m.Unlock() <- This fixes the issue
}()
go func() {
    time.Sleep(1)
    m.Lock()
    defer m.Unlock()
    done <- true
}()
<-done
```

### Example 2.3: The hungry philosophers problem

Locking mutexes in the wrong order will often result in a deadlock. The fix is always to lock them in the same order everywhere.

```go
var m1, m2 sync.Mutex

done := make(chan bool)

go func() {
    m1.Lock(); defer m1.Unlock()
    time.Sleep(1)
    m2.Lock(); defer m2.Unlock()
    done <- true
}()

go func() {
    m2.Lock(); defer m2.Unlock()
    time.Sleep(1)
    m1.Lock(); defer m1.Unlock()
    done <- true
}()
<-done; <-done
```

## Gotcha 3: Goroutine leaks

In this example, a timeout can lead to a goroutine leak.

The goroutine is launched with an unbuffered channel, and the select statement will wait for the goroutine to complete or for the timeout to be reached. If the timeout is reached before the goroutine completes, the goroutine will block forever, and the goroutine will never be garbage collected.

```go
func finishReq(timeout time.Duration) *obj {
    ch := make(chan obj)
    go func() {
        . . .       // work that takes too long
        ch <- fn()  // blocking send
    }()

    select {
    case rslt := <-ch:
        return rslt
    case <-time.After(timeout):
        return nil
    }
}
```

One possible solution is to use a buffered channel to avoid blocking the goroutine if the timeout is reached before the goroutine completes.

## Gotcha 4: Incorrect use of WaitGroup

In this example we are adding to the WaitGroup after starting the unit of work.
This is wrong as there is a chance the following can happen:

1. `walkdir` is accessed the first time.
2. The WaitGroup (`wg`) is incremented, and a `wg.Done` is defered.
3. The visit function is called, but before `walkDir` gets to the point of adding to the WaitGroup, the function returns.
4. This makes `wg` to be ready before the work is done.

```go
func walkDir(dir string, pairs chan<- pair, ...) {
    . . . // Something that takes a lot of time
    wg.Add(1)
    defer wg.Done()
    visit := func(p string, fi os.FileInfo, ...) {
        if fi.Mode().IsDir() && p != dir {
            go walkDir(p, pairs, wg, limits)
            . . .
        }
    }
}
err := walkDir(dir, paths, wg)
wg.Wait()
```

A solution would be to make sure that we always add to the WaitGroup before starting the unit of work.

```go
func walkDir(dir string, pairs chan<- pair, ...) {
    . . . // Something that takes a lot of time
    defer wg.Done()
    visit := func(p string, fi os.FileInfo, ...) {
        if fi.Mode().IsDir() && p != dir {
            wg.Add(1)
            go walkDir(p, pairs, wg, limits)
            . . .
        }
    }
}
wg.Add(1)
err := walkDir(dir, paths, wg)
wg.Wait()
```

## Gotcha 5: Closure capture

A goroutine closure shouldn’t capture a mutating variable

```go
// Bad example:
for i := 0; i < 10; i++ {
    go func() {
        fmt.Println(i)
    }()
}
```

Instead, pass the variable’s value as a parameter. (Passing by value)

```go
// Good example 1:
for i := 0; i < 10; i++ {
    go func(i int) {
        fmt.Println(i)
    }(i)
}
```

Another solution would be to create a new variable inside the loop.

```go
// Good example 2:
for i := 0; i < 10; i++ { // RIGHT
    i := i
    go func() {
        fmt.Println(i)
    }()
}
```

## Gotcha 6: Select can lead to mistakes

### Example 6.1: Skipping a channel due to a default case

The `default` case is always active. If nothing else matches, the default case is selected.

```go
for {
    x := socket.Read()
    select {
    case output <- x:
    . . .
    default:
        return
    }
}
```

### Example 6.2: Selecting a nil channel

Given that a `nil` channel is always ignored, the following code will never return.

```go
var output chan int // nil channel
for {
    x := socket.Read()
    select {
    case output <- x:
        // This case will never be selected because output is nil
    default:
        // Handle the case where the channel is nil
    }
}

```

### Example 6.3: Sending to a channel that is full is skipped

A full channel (for send) is skipped over.

```go
package main

import "fmt"

func main() {
    output := make(chan int, 2) // buffered channel with capacity 2

    // Fill the channel to its capacity
    output <- 1
    output <- 2

    for i := 3; i <= 5; i++ {
        select {
        case output <- i:
            fmt.Println("Sent:", i)
        default:
            fmt.Println("Channel is full, skipping:", i)
        }
    }
}

```

### Example 6.4: Beware of using done channels to exit goroutines

```go
for {
    select {
    case x := <- input:
        . . .
    case <- done:
        return
    }
}
```

In the above code, the select statement is used to read from either the input channel or the done channel. However, the select statement does not prioritize one case over the other. This means that if a value is available on both channels, it is non-deterministic which case will be executed. As a result, the loop might exit before all values from the input channel are processed if the done channel is selected.

To ensure that all values from the input channel are read before exiting, use the done channel only for error handling or aborting the operation. Close the input channel when there are no more values to read (e.g., on EOF).

```go
go func() {
    // Simulate reading input and closing the channel on EOF
    for _, value := range data {
        input <- value
    }
    close(input)
}()

for {
    select {
    case x, ok := <-input:
        if !ok {
            // input channel is closed, exit the loop
            return
        }
        // Process the value x
        . . .
    case <-done:
        // Handle error or abort
        return
    }
}
```

In this improved approach, the input channel is closed when there are no more values to read. The select statement now checks if the input channel is closed by using the ok variable. If ok is false, it means the channel is closed, and the loop exits. This ensures that all values from the input channel are processed before exiting. The done channel is still used for error handling or aborting the operation.

### Available channels are selected at random!

Available channels are selected at random. If multiple channels are ready at the same time, one is chosen at random. Values are not considered from top to bottom.

## Four considerations when using concurrency

1. Don't start a goroutine without knowing how it will stop
2. Acquire locks/semaphores as late as possible; release them in the reverse order
3. Don't wait for non-parallel work that you could do yourself

```go
func do() int {
    ch := make(chan int)
    go func() { ch <- 1 }()
    return <-ch
}
```

4. Simplify! Review! Test!
