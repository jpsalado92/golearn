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
