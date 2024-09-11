# [Go Class: 30 Concurrency Gotchas](https://www.youtube.com/watch?v=K1hwpNnCJgY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=31)

- [Go Class: 30 Concurrency Gotchas](#go-class-30-concurrency-gotchas)
  - [Gotcha 1: Race conditions](#gotcha-1-race-conditions)
  - [Gotcha 2: Dead locks](#gotcha-2-dead-locks)
    - [Description](#description)
    - [Prevention](#prevention)
    - [Example 1](#example-1)
    - [Example 2](#example-2)
    - [Example 3](#example-3)
  - [Gotcha 3: Goroutine leaks](#gotcha-3-goroutine-leaks)
    - [Description](#description-1)
    - [Prevention](#prevention-1)
    - [Example 1](#example-1-1)
  - [Gotcha 4: Incorrect use of WaitGroup](#gotcha-4-incorrect-use-of-waitgroup)
    - [Description](#description-2)
    - [Prevention](#prevention-2)
    - [Example](#example)
  - [Gotcha 5: Closure capture](#gotcha-5-closure-capture)
  - [Gotcha 6: Select can lead to mistakes](#gotcha-6-select-can-lead-to-mistakes)
    - [Mistake #1: skipping a full channel to default and losing a message](#mistake-1-skipping-a-full-channel-to-default-and-losing-a-message)
  - [Mistake #2: reading a “done” channel and aborting when input is backed up on another channel — that input is lost](#mistake-2-reading-a-done-channel-and-aborting-when-input-is-backed-up-on-another-channel--that-input-is-lost)
    - [Four considerations when using concurrency:](#four-considerations-when-using-concurrency)

## Gotcha 1: Race conditions
A race condition occurs when the outcome of a program depends on the timing or interleaving of multiple threads or processes accessing a shared resource.

In simpler terms, it's like multiple people trying to edit a document at the same time without coordinating. The final version might be a mishmash of changes, leading to unexpected or incorrect results. This can happen in software when multiple threads or processes are trying to modify the same data simultaneously, and the outcome depends on the order in which they do so.### Prevention

In the following example, requests directed to `/` lead to the handler function printing the value of `nextID` and incrementing it.

The race condition occurs when multiple requests are made simultaneously, and the value of `nextID` is read and incremented by multiple goroutines at the same time. This can lead to unexpected results, such as two goroutines reading the same value of `nextID` and incrementing it, resulting in the same value being returned for both requests.

```go
var nextID = 0

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>You got %v<h1>", nextID)
    nextID++
}

func main() {
    http.HandleFunc("/", handler)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
```
## Gotcha 2: Dead locks
### Description
### Prevention
### Example 1
Go can usually detect when no goroutine is able to make progress; here the main goroutine is blocked on a channel it can never read.
```go
func main() {
    ch := make(chan bool)
    go func(ok bool) {
        fmt.Println("STARTED")
        if ok {
            ch <- ok
        }
        }(false)
    <-ch
    fmt.Println("DONE")
}
```
### Example 2
Locking a mutex and then failing to unlock it afterwards; the fix is to use defer at the point of locking.
```go
var m sync.Mutex
    done := make(chan bool)
    go func() {
        m.Lock() // not unlocked!
    }()
    go func() {
        time.Sleep(1)
        m.Lock()
        defer m.Unlock()
        done <- true
    }()
    <-done
```
### Example 3
Locking mutexes in the wrong order will often result in deadlock; the fix is always to lock them in the same order everywhere.
This example is also refered as the hungry philosopher problem.
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
### Description
### Prevention
### Example 1
In this example, a timeout can lead to a goroutine leak.

The goroutine is launched with an unbuffered channel, and the select statement will wait for the goroutine to complete or for the timeout to be reached. If the timeout is reached before the goroutine completes, the goroutine will block forever, and the goroutine will never be garbage collected.

The solution is to use a buffered channel or to use a context to cancel the goroutine.

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
## Gotcha 4: Incorrect use of WaitGroup
### Description
### Prevention
### Example
In this example we are adding to the WaitGroup after starting the unit of work.
This is wrong as there is a chance the following can happen:
1. Walkdir is accessed the first time.
2. The WaitGroup is incremented, and a Done is defered.
3. The visit function is called, but before walkDir gets to the point of adding to the WaitGroup, the function returns.
4. This makes the Wait group to be ready before the work is done.

```go
func walkDir(dir string, pairs chan<- pair, ...) {
    wg.Add(1) // WRONG
    defer wg.Done()

    visit := func(p string, fi os.FileInfo, ...) {
        if fi.Mode().IsDir() && p != dir {
            // wg.Add(1) // RIGHT
            go walkDir(p, pairs, wg, limits)
            . . .
        }
    }
}

// wg.Add(1) // RIGHT
err := walkDir(dir, paths, wg)

wg.Wait()
```
We should always add to the WaitGroup before starting the unit of work.
## Gotcha 5: Closure capture

A goroutine closure shouldn’t capture a mutating variable
```go
for i := 0; i < 10; i++ { // WRONG
    go func() {
        fmt.Println(i)
    }()
}
```
Instead, pass the variable’s value as a parameter. (Passing by value)
```go
for i := 0; i < 10; i++ { // RIGHT
    go func(i int) {
        fmt.Println(i)
    }(i)
}
```

Another solution would be to create a new variable inside the loop.
```go
for i := 0; i < 10; i++ { // RIGHT
    i := i
    go func() {
        fmt.Println(i)
    }()
}
```
## Gotcha 6: Select can lead to mistakes

- `default` is always active. If nothing else matches, the default case is selected.
- A `nil` channel is always ignored.
- A full channel (for send) is skipped over.
- A “done” channel is just another channel. The fact that it’s called “done” doesn’t make it special.
- Available channels are selected at random. If multiple channels are ready at the same time, one is chosen at random. Values are not considered from top to bottom.


### Mistake #1: skipping a full channel to default and losing a message
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

The code was written assuming we'd skip output only if it was set to nil
We also skip if output is full, and lose this and future messages

## Mistake #2: reading a “done” channel and aborting when input is backed up on another channel — that input is lost

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

There’s no guarantee we read all of input before reading done
Better: use done only for an error abort; close input on EOF

### Four considerations when using concurrency:
1. Don't start a goroutine without knowing how it will stop
2. Acquire locks/semaphores as late as possible; release them in the reverse order
3. Don’t wait for non-parallel work that you could do yourself
```go
func do() int {
    ch := make(chan int)
    go func() { ch <- 1 }()
    return <-ch
}
```
4. Simplify! Review! Test!