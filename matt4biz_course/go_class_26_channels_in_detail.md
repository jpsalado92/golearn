[Go Class: 26 Channels in Detail](https://www.youtube.com/watch?v=fCkxKGd6CVQ&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=27)

- [Channel block behavior](#channel-block-behavior)
- [Blocking buffer vs unbuffered channel](#blocking-buffer-vs-unbuffered-channel)
- [Closed channels](#closed-channels)
- [Nil channels](#nil-channels)
- [Reasons for buffering](#reasons-for-buffering)
- [Counting semaphore](#counting-semaphore)

## Channel block behavior

A channel can be **read** from (`<-chan`) if:

- At least one writer is ready to write (rendezvous)
- There is unread data in its buffer
- It is closed

A channel can be **written** to (`chan<-`) if:

- At least one reader is ready to read (rendezvous)
- It has buffer space

## Blocking buffer vs unbuffered channel

In the example below, the channel is unbuffered, so the write operation is blocked waiting for a reader, which is never reached as it is in the next line.

```go
ch := make(chan int)
ch <- 1
a:= <-ch
```

In the example below, the channel is buffered, so the write operation is not blocked.

```go
ch := make(chan int, 1)
ch <- 1
a:= <-ch
```

## Closed channels

Closing a channel is a way to tell that some work is done.

Closing a channel is a way to prevent goroutines from blocking indefinitely.

```go

// Function that writes to a channel and leaks
func write(ch chan int) {
    for i := 0; i++ {
        ch <- i
    }
}
// Read from a channel twice
ch := make(chan int)
go write(ch)
a := <-ch
b := <-ch

close(ch)  // Close the channel, preventing the write goroutine from blocking indefinitely
```

Closed channels return zero values for reads and writes. One can tell if a channel is closed by using the second return value of a read operation.

```go
close(ch)
v, ok := <-ch  // 0, false
```

A channel can only be closed once, subsequent closes will panic. And only one goroutine should close a channel.

## Nil channels

We can suspend channels by setting them to nil.

If reading or writing to a nil channel, the program will block,
but in a `select` statement it will ignore the channel.

## Reasons for buffering

1. Avoid goroutine leaks (from an abandoned channel)
2. Avoid rendezvous pauses (performance improvement)

Do not buffer until its needed, buffering may hide a race condition.
When using buffers, setting the proper buffer size is a matter of testing with each use case.

## Counting semaphore

A counting semaphore limits work in progress (or occupancy) to a certain number.
Once it's full, only one unit of work can enter for each unit that leaves.

This can be modeled through a buffered channel.

- Attempt to send (write) before starting work.
- The send will block if the buffer is full.
- Receive (read) after work is done to free up space, allowing another unit of work to start.
