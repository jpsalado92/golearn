[Go Class: 28 Conventional Synchronization](https://www.youtube.com/watch?v=DtXNSE3Yejg&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=29)

- [Best practices](#best-practices)
- [Tool `sync.Once`](#tool-synconce)
- [Tool `sync.Pool`](#tool-syncpool)
- [Race condition case](#race-condition-case)
	- [Example 1.1: Broken code with race condition](#example-11-broken-code-with-race-condition)
	- [Example 1.2: Semaphore fix](#example-12-semaphore-fix)
	- [Example 1.3: Mutex fix](#example-13-mutex-fix)
	- [Example 1.4: Atomic fix](#example-14-atomic-fix)

## Best practices

- Embed the `sync.Mutex` in the struct. Use Lock and Unlock methods to protect the struct fields.
- Use `sync.RWMutex` for read-write locks. Read locks can be acquired by multiple goroutines. Use `RLock` and `RUnlock` for them.

## Tool `sync.Once`

A `sync.Once` is used to run a function exactly once. It is useful for lazy initialization.

```go
var once sync.Once
var x *singleton

func initialize() {
	x = new(singleton)
}

Pfunc handle(w http.ResponseWriter, r *http.Request) {
	once.Do(initialize) // This is safer than checking for nil and then initializing
}
```

## Tool `sync.Pool`

A `Pool` provides for efficient & safe reuse of objects. A `sync.Pool` is often used to cache objects that are expensive to create. It is also useful for reducing the number of allocations and garbage collection.

A `Pool` has two methods:

- `Get` returns an object from the pool. If the pool is empty, it creates a new object.
- `Put` puts an object back into the pool.

sync.Pool is concurrent-safe. It is safe to use it from multiple goroutines. The only limitation is that the pool itself sees objects stores within it as empty interfaces. So, you need to use reflection to convert them back to the original type.

```go
var pool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)  // Here we are returning interface{} type
	},
}

func log(s string) {
	b := pool.Get().(*bytes.Buffer) // Reflection is used to convert the interface{} to *bytes.Buffer
	b.Reset()
	b.WriteString(s)
	fmt.Println(b.String())
	pool.Put(b)
}
```

## Race condition case

### Example 1.1: Broken code with race condition

```go

package main

import (
	"fmt"
	"sync"
)

func do() int {
    var x int64
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            x++
            wg.Done()
        }()
    }
    wg.Wait()
    return int(x)
}

func main() {
    fmt.Println(do())
}
```

### Example 1.2: Semaphore fix

```go
package main

import (
	"fmt"
	"sync"
)

func do() int {
	var x int64
	m := make(chan bool, 1)
	wg := sync.WaitGroup{}
	m <- true
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			<-m
			x++
			m <- true
			wg.Done()
		}()
	}
	wg.Wait()
	return int(x)
}

func main() {
	fmt.Println(do())
}
```

### Example 1.3: Mutex fix

```go
package main

import (
	"fmt"
	"sync"
)

func do() int {
	var x int64
	var m sync.Mutex
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			x++
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return int(x)
}

func main() {
	fmt.Println(do())
}
```

### Example 1.4: Atomic fix

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func do() int {
	var x int64
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			atomic.AddInt64(&x, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	return int(x)
}

func main() {
	fmt.Println(do())
}
```
