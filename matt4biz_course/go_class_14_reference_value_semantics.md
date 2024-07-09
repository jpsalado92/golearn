# [Go Class: 14 Reference & Value Semantics](https://www.youtube.com/watch?v=904pyovPvXM&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=15)

- [Go Class: 14 Reference \& Value Semantics](#go-class-14-reference--value-semantics)
  - [Reference \& Value Semantics](#reference--value-semantics)
  - [When to use pointers](#when-to-use-pointers)
  - [Gotchas](#gotchas)
    - [Gotcha 1: Loops](#gotcha-1-loops)
    - [Gotcha 2: Always return the slice](#gotcha-2-always-return-the-slice)
    - [Gotcha 3: Do not keep pointers to slice values](#gotcha-3-do-not-keep-pointers-to-slice-values)
    - [Gotcha 4: Example of messed up loop](#gotcha-4-example-of-messed-up-loop)

## Reference & Value Semantics
Passing down pointers (*Reference Semantincs*) means that: We want to share resources, not copy them.

Passing down values (*Value Semantincs*) means that: We want to make a copy of the resource, not share it.


## When to use pointers

- When objects are not copyable, mutex.
- When using wait groups.
- When objects are too large in size (> 64 bytes).
- When objects are expected to be mutated later.
- When objects are intended to be filled later. (JSON unmarshalling)
- When we want to refer to "null" objects.


## Gotchas

### Gotcha 1: Loops

When iterating a slice of objects that we want to mutate, we should use the index to access the object that we want to mutate. Using the object directly through the range will create a copy of the object, and we will not be able to mutate the original object.

```go
for i, thing:= range things {
    // Thing is a copy
    ...
}

for i := range things {
    things[i].doSomething() // This is a reference, so we can mutate the original object through the index
    ...
}
```

### Gotcha 2: Always return the slice

When using functions that mutate slices, we should always return back the slice. This is because appending to a slice can reallocate the underlying data.
```go
func appendToSlice(slice []int, value int) []int {
    slice = append(slice, value)
    return slice
}

func main() {
    slice := []int{1, 2, 3}
    slice = appendToSlice(slice, 4)
}
```


### Gotcha 3: Do not keep pointers to slice values

When we append to a slice, the underlying array can be reallocated. If we keep a pointer to the slice, the pointer can become invalid.


### Gotcha 4: Example of messed up loop

```go
package main

import "fmt"

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
    a:= [][]byte{}

    for _, item := range items {
        a = append(a, item[:])
    }

	fmt.Println(items)
	fmt.Println(a)
}
```
```
[[1 2] [3 4] [5 6]]
[[5 6] [5 6] [5 6]]
```

This happened because in the loop, when we append `item[:]` to the slice, we are appending a reference to a memory address that is being reused. So, all the elements in the slice are pointing to the same memory address which, in the end, contains the last value of the original slice.

In order to fix it:
```go
package main

import "fmt"

func main() {
    items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
    a:= [][]byte{}

    for _, item := range items {
        itemCopy := make([]byte, len(item))
        copy(itemCopy, item[:])
        a = append(a, itemCopy)
    }

    fmt.Println(items)
    fmt.Println(a)
}
```
```
[[1 2] [3 4] [5 6]]
[[1 2] [3 4] [5 6]]
```

