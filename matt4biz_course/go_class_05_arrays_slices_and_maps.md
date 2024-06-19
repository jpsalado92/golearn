# [Go Class: 05 Arrays, Slices, and Maps](https://www.youtube.com/watch?v=T0Xymg0_aSU&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=6)
- [Go Class: 05 Arrays, Slices, and Maps](#go-class-05-arrays-slices-and-maps)
  - [Arrays](#arrays)
  - [Slices](#slices)
  - [Main differences table](#main-differences-table)
  - [Maps](#maps)
    - [How to tell if a key is in a map:](#how-to-tell-if-a-key-is-in-a-map)

## Arrays
- Sequence of elements.
- Arrays are fixed size.
- Length is defined on declaration.
- Not used as much in Go.
```go
// All these are equivalent
var a [3]int
var a [3]int{0, 0, 0}
var a [...]int{0, 0, 0}
```

When assigned to a new array, it creates a copy of the array, not a reference.
```go
a := [...]int{1, 2, 3}
b := a
fmt.Printf("a: %p\n", &a) // Points to first element of a, which is in X
fmt.Printf("b: %p\n", &b) // Points to first element of b, which is in Y
```

Cannot assign arrays of different sizes.
```go
a := [3]int{1, 2, 3}
b := [4]int{1, 2, 3, 4}
a = b // Error, type mismatch
```

## Slices
- Sequence of elements.
- A.k.a variable length array.
- Length is not defined on declaration.
```go
var a []int  // Slice of int, nil
var a []int{1, 2 ,3} // Slice of int, length 3
a := make([]int, 3) // Slice of int, length 3
```
When assigned to a new slice, it creates a reference to the original slice.
```go
a := []int{1, 2, 3}
b := a
fmt.Printf("a: %p\n", &a) // Points to first element of a, which is in X
fmt.Printf("b: %p\n", &b) // Points to first element of a, which is in X
```

In the examples above, `a` and `b` are slice descriptors. They contain a pointer to the first element of the slice, the length of the slice, and the capacity of the slice.
```go
// Appending
a := []int{1, 2, 3}
a = append(a, 4)
fmt.Println(a) // [1 2 3 4]
```

If the capacity of the slice is exceeded (for example, by appending new elements to the slcie), a new slice is created with double the capacity somewhere else in memory, and the elements are copied over. This is done to prevent the need to copy elements over every time an element is added.

Indexing a slice returns the element at that index.
```go
a := []int{1, 2, 3}
fmt.Println(a[0]) // 1
```

Slicing a slice returns a new slice with the elements from the start index to the end index.
```go
a := []int{1, 2, 3, 4, 5}
fmt.Println(a[1:3]) // [2 3]
```

## Main differences table
| Array | Slice |
|-------|-------|
| Fixed size | Variable size |
| Length defined on declaration | Length not defined on declaration |
| Copy on assignment | Reference on assignment |
| Cannot assign arrays of different sizes | Can assign slices of different sizes |
| Can be used as a key in a map | Cannot be used as a key in a map |
| Has Copy and Append methods |  |
| Comparable | Not comparable |
| PseudoConstant |  |
## Maps
- Map of keys to values.
- Can read from nil map, but cannot write to it.
- Maps are passed by reference.
- They key used in a map must be comparable.
- Maps can only be compared to `nil`.
```go
var a map[string]int  // key is string, value is int
```
```go
var m map[string]int  // nil, no storage
p := make(map[string]int)  // no-nil, empty
a := p["the"]   // 0, key not found
b := m["the"]  // 0, key not found

m["the"] = 1 // PANIC - assignment to entry in nil map
m = p
m["and"]++ // OK
c := p["and"] // returns 1
```

### How to tell if a key is in a map:
```go
m := map[string]int{"the": 1, "and": 1}
v, ok := m["the"]
if ok {
    fmt.Println("The key is in the map")
}
```