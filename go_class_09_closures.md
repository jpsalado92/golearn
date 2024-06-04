# [Go Class: 09 Closures](https://www.youtube.com/watch?v=US3TGA-Dpqo&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=11)

- [Go Class: 09 Closures](#go-class-09-closures)
  - [Scope vs Lifetime](#scope-vs-lifetime)
  - [What is a closure?](#what-is-a-closure)
    - [Example 1: Fibonacci closure](#example-1-fibonacci-closure)
    - [Example 2: Closures do not share resources](#example-2-closures-do-not-share-resources)


## Scope vs Lifetime

- Scope is static, based on the code at compile time.
- Lifetime depends on runtime.

In this example, `doIt` is returning a pointer to a local variable. So the value will live as long as the program keeps a pointer to it, even in the main function.
```go
package main

func doIt() *int {
    b := 10
    return &b  // By this moment, Go knows that this will have to live more than the function call.
}

func main() {
    doIt()

}
```

## What is a closure?

A closure is a function value that refers to variables outside of the scope of its own function body.

A closure can survive the scope of the function that defines it. This means it has access to variables that are outside of its scope.


### Example 1: Fibonacci closure

```go
package main

import "fmt"

func fib() func() int {  // This function returns a function that returns an int
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return b
    }
}

func main() {
    f := fib()  // fib() will be returning an anonymous function, but its actually a closure
                // As long as f exist, a and b will exist.
                // f refers to a closure
                // fib() is a closure

for x := f(); x < 100; x = f() { // Every time f() is called, the variables a and b are updated.
        fmt.Println(x)
    }
}
```

### Example 2: Closures do not share resources

Every call to fib() will return a new closure. So, if we call fib() twice, we will have two different closures.

```go
package main

import "fmt"

func fib() func() int {  // This function returns a function that returns an int
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return b
    }
}

func main() {
    f, g := fib(), fib()
    fmt.Println(f(), f(), f(), f(), f()) // 1 2 3 5 8
    fmt.Println(g(), g(), g(), g(), g()) // 1 2 3 5 8
}
```

