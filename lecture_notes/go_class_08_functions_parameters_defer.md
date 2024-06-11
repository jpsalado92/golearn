# [Go Class: 08 Functions, Parameters & Defer](https://www.youtube.com/watch?v=wj0hUjRHkPs&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=10)
- [Go Class: 08 Functions, Parameters \& Defer](#go-class-08-functions-parameters--defer)
  - [Functions](#functions)
    - [Parameters](#parameters)
      - [Example 1: Array passed down to a function by value](#example-1-array-passed-down-to-a-function-by-value)
      - [Example 2: Slice passed down to a function by reference](#example-2-slice-passed-down-to-a-function-by-reference)
      - [Example 3: Map passed down to a function by reference](#example-3-map-passed-down-to-a-function-by-reference)
      - [Example 4: Overwriting the original reference by using a pointer](#example-4-overwriting-the-original-reference-by-using-a-pointer)
    - [Return values](#return-values)
  - [Defer](#defer)
    - [Example 1: Hello World](#example-1-hello-world)
    - [Example 2: Defer for closing a file](#example-2-defer-for-closing-a-file)
    - [Example 3: Stacked defers](#example-3-stacked-defers)
    - [Defer Gotcha #1](#defer-gotcha-1)
    - [Defer Gotcha #2](#defer-gotcha-2)


## Functions

The **signature** of a function in go is the order and type of the parameters and the return values. It does not depend on the name of the function or its parameters.

### Parameters

When declaring a function, you might do it with **formal parameters**, but when calling it, you might use **actual parameters** (also known as arguments).

Parameters in Go can be passed either by value or by reference.

**By Value**: The function receives a copy of the parameter. This applies to:
- Numbers
- Booleans
- Structs
- Arrays

**By Reference**: The function receives a reference to the parameter. This applies to:
- Slices
- Maps
- Channels
- Pointers
- Interfaces
- Strings




#### Example 1: Array passed down to a function by value
```go

package main

import (
	"fmt"
)

func do(my_array [3]int) [3]int {
    my_array[0] = 0
    return my_array
}

func main() {
    my_array := [3]int{1, 2, 3}
    new_array := do(my_array)

    fmt.Println(new_array)  // [0 2 3]
    fmt.Println(my_array)  // [1 2 3]
}

```

#### Example 2: Slice passed down to a function by reference

Pointers pointing to the same memory address

```go
package main

import (
	"fmt"
)

func do(my_slice []int) []int {
    my_slice[0] = 0
    fmt.Printf("my_slice@ %p\n", my_slice) // my_array@ 0xc0000b6010
    return my_slice
}

func main() {
    my_slice := []int{1, 2, 3}
    fmt.Printf("my_array@ %p\n", my_slice) // my_array@ 0xc0000b6010
    same_array := do(my_slice)
    fmt.Println(my_slice)  // [0 2 3]
    fmt.Println(same_slice)  // [0 2 3]
}

```

#### Example 3: Map passed down to a function by reference

Here a map is passed by reference, but the function creates a new map, so the original map is not modified.

```go
package main

import (
	"fmt"
)

func do(my_map map[string]int) map[string]int {
    my_map["a"] = 0
    fmt.Printf("my_map@ %p\n", my_map) // my_map@ 0xc0000b6010
    my_map = make(map[string]int)
    fmt.Printf("my_map@ %p\n", my_map) // my_map@ 0xc0000b7000
    return my_map
}

func main() {
    my_map := map[string]int{"a": 1, "b": 2, "c": 3}
    fmt.Printf("my_map@ %p\n", my_map)  // my_map@ 0xc0000b6010
    new_map := do(my_map)
    fmt.Println(my_map) // map[a:0 b:2 c:3]
    fmt.Println(new_map) // map[]
}

```

#### Example 4: Overwriting the original reference by using a pointer

In this example, we pass a pointer to the map, so the function can modify the original map.

When the overwrite occurs, the original map is modified.

```go
package main

import (
	"fmt"
)

func do(fmap *map[string]int) { // Here the asterisk is used to denote that the argument is a pointer to a map.
	(*fmap)["a"] = 0 // We use *fmap to access the value of the map.
	fmt.Println("fmap", fmap)
	*fmap = make(map[string]int) // Refering to the original map with the pointer.
    (*fmap)["z"] = 100
	fmt.Println("fmap", *fmap)
}

func main() {
	my_map := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("my_map", my_map)
	do(&my_map) // Passing the address of the map to the function.
	fmt.Println("my_map", my_map)
}

```

### Return values

Return values are written with the following syntax:

```go

// func <function_name>(<parameters>) <return_type> {
//     return <value>
```

The return values can be named (`naked`), so they can be returned without explicitly writing them.
```go
func do() (a int, b int) {
    a = 1
    b = 2
    return
}

func main() {
    x, y := do()
    fmt.Println(x, y)
}
```

```go
// Errors are commonly returned as the last value of a function.
func do3() (int, error) {
    return 1, nil
}
```

## Defer

The `defer` statement captures a function *call* to run after the function that contains it finishes.

This is useful for cleaning up resources, closing files, etc.

### Example 1: Hello World

```go

package main

import (
    "fmt"
)

func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}

```

The output of this code will be:
```
Hello
World
```

The `defer` statement is executed after the `fmt.Println("Hello")` statement, but before the `main` function ends.

### Example 2: Defer for closing a file

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()  // It runs after we know the file was opened.
    fmt.Fprintln(f, "Hello, World!")
}
```

In this example, the `defer` statement is used to close the file after it has been opened and written to.

### Example 3: Stacked defers
Defer statements work as a stack, LIFO (Last In, First Out).
```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Counting")
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("Done")
}
```

The output of this code will be:
```
Counting
Done

2
1
0
```

The `defer` statement captures the value of `i` at the time it is called, so it will print the values in reverse order.
Note that the arguments to defer are copied on call, so the value of `i` is copied when the defer is called, not when it is actually executed.


### Defer Gotcha #1

A `defer` statement captures the value of the variables at the time it is called, not when it is executed.

```go
package main

import (
    "fmt"
)

func main() {
    a := 1
    defer fmt.Println(a)
    a = 2
}
```

The output of this code will be `1` not `2`.

### Defer Gotcha #2

A `defer` statement runs before a `return` statement, so it can affect the output of the function.

```go
package main

func do() (a int) {
    defer func() {
        a = 2
    }()
    a = 1
    return a
}

func main() {
    a := do()
    println(a)  // The printed value will be 2, not 1.
}
```


