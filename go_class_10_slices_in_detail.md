# [Go Class: 10 Slices in Detail](https://www.youtube.com/watch?v=pHl9r3B2DFI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=12)

- [Go Class: 10 Slices in Detail](#go-class-10-slices-in-detail)
  - [Slice types](#slice-types)
    - [Nil Slice](#nil-slice)
    - [Empty Slice](#empty-slice)
    - [Empty slice with only length defined](#empty-slice-with-only-length-defined)
    - [Slice with length and capacity defined](#slice-with-length-and-capacity-defined)
    - [Summary Table](#summary-table)
  - [Empty vs Nil](#empty-vs-nil)
    - [GOTCHA: Differences on json encoding](#gotcha-differences-on-json-encoding)
    - [TIP: Avoid checking for nil before checking emptiness, instead directly use len()](#tip-avoid-checking-for-nil-before-checking-emptiness-instead-directly-use-len)
  - [Overflowing slice capacity](#overflowing-slice-capacity)
    - [Overflow memory reallocation](#overflow-memory-reallocation)
  - [Inheriting capacity when slicing a slice](#inheriting-capacity-when-slicing-a-slice)

## Slice types
### Nil Slice

```go	
package main

import "fmt"

func main() {
    var s []int  // NIL slice definition

    fmt.Printf("Slice: %v\n", s)
    fmt.Printf("Type: %T\n", s)
    fmt.Printf("Length: %d\n", len(s))
    fmt.Printf("Capacity: %d\n", cap(s))
    fmt.Printf("Is slice nil?: %t\n", s == nil)
    fmt.Printf("Slice representation: %#v\n", s)
}
```
```
Slice: []
Type: []int
Length: 0
Capacity: 0
Is slice nil?: true
Slice representation: []int(nil)
```

### Empty Slice

```go	
package main

import "fmt"

func main() {
    t := []int{}  // Empty slice definition

    fmt.Printf("Slice: %v\n", t)
    fmt.Printf("Type: %T\n", t)
    fmt.Printf("Length: %d\n", len(t))
    fmt.Printf("Capacity: %d\n", cap(t))
    fmt.Printf("Is slice nil?: %t\n", t == nil)
    fmt.Printf("Slice representation: %#v\n", t)
}
```
```
Slice: []
Type: []int
Length: 0
Capacity: 0
Is slice nil?: false
Slice representation: []int{}
```

### Empty slice with only length defined

```go
package main

import "fmt"

func main() {
    u := make([]int, 5)

    fmt.Printf("Slice: %v\n", u)
    fmt.Printf("Type: %T\n", u)
    fmt.Printf("Length: %d\n", len(u))
    fmt.Printf("Capacity: %d\n", cap(u))
    fmt.Printf("Is slice nil?: %t\n", u == nil)
    fmt.Printf("Slice representation: %#v\n", u)
}
```
```
Slice: [0 0 0 0 0]
Type: []int
Length: 5
Capacity: 5
Is slice nil?: false
Slice representation: []int{0, 0, 0, 0, 0}
```

### Slice with length and capacity defined

```go
package main

import "fmt"

func main() {
    v := make([]int, 0, 5)

    fmt.Printf("Slice: %v\n", v)
    fmt.Printf("Type: %T\n", v)
    fmt.Printf("Length: %d\n", len(v))
    fmt.Printf("Capacity: %d\n", cap(v))
    fmt.Printf("Is slice nil?: %t\n", v == nil)
    fmt.Printf("Slice representation: %#v\n", v)
}
```
```
Slice: []
Type: []int
Length: 0
Capacity: 5
Is slice nil?: false
Slice representation: []int{}
```

### Summary Table
| Case                           | Definition             | Value       | Type  | Length | Capacity | Is nil? | Representation       |
| ------------------------------ | ---------------------- | ----------- | ----- | ------ | -------- | ------- | -------------------- |
| Nil Slice                      | var s []int            | []          | []int | 0      | 0        | true    | []int(nil)           |
| Empty Slice                    | t := []int{}           | []          | []int | 0      | 0        | false   | []int{}              |
| Slice with length and capacity | u := make([]int, 5)    | [0 0 0 0 0] | []int | 5      | 5        | false   | []int{0, 0, 0, 0, 0} |
| Slice with only capacity       | v := make([]int, 0, 5) | []          | []int | 0      | 5        | false   | []int{}              |

## Empty vs Nil

### GOTCHA: Differences on json encoding
    
```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var a []int  // NIL slice definition
    b := []int{}  // Empty slice definition

    sJson, _ := json.Marshal(a)
    tJson, _ := json.Marshal(b)

    fmt.Printf("Nil slice: %s\n", sJson)
    fmt.Printf("Empty slice: %s\n", tJson)
}
```
```
Nil slice: null
Empty slice: []
```

### TIP: Avoid checking for nil before checking emptiness, instead directly use len()

```go
package main

import "fmt"

func main() {
    var s []int  // NIL slice definition
    t := []int{}  // Empty slice definition

    fmt.Printf("Is nil slice empty?: %t\n", len(s) == 0)
    fmt.Printf("Is empty slice empty?: %t\n", len(t) == 0)
}
```
``` 
Is nil slice empty?: true
Is empty slice empty?: true
```

## Overflowing slice capacity


```go
package main

import "fmt"

func main() {
	v := make([]int, 0, 5) // len(v)=0, cap(v)=5
	v = append(v, 1, 2, 3, 4, 5) // len(v)=5, cap(v)=5
	fmt.Printf("len=%d cap=%d %v\n", len(v), cap(v), v)
	v = append(v, 6)       // len(v)=6, cap(v)=10
	fmt.Printf("len=%d cap=%d %v\n", len(v), cap(v), v)
}
```
```
len=5 cap=5 [1 2 3 4 5]
len=6 cap=10 [1 2 3 4 5 6]
```

As seen in the example, whenever the capacity of a slice is exceeded, the slice is reallocated with a new capacity. The new capacity is calculated as 2 times the old capacity. This is done to avoid frequent reallocations and copying of the slice elements.

### Overflow memory reallocation

```go
package main

import "fmt"

func main() {
    v := make([]int, 5) // len(v)=5, cap(v)=5
    fmt.Printf("len=%d cap=%d %v\n", len(v), cap(v), v)
    fmt.Printf("Memory address for v: %p\n", &v)
    fmt.Printf("Memory address for v[0]: %p\n", &v[0])
    
    v = append(v, 0) // len(v)=6, cap(v)=10
    fmt.Printf("len=%d cap=%d %v\n", len(v), cap(v), v)
    fmt.Printf("Memory address for v: %p\n", &v)
    fmt.Printf("Memory address for v[0]: %p\n", &v[0])
}
```
```
len=5 cap=5 [0 0 0 0 0]
Memory address for v: 0xc000010018
Memory address for v[0]: 0xc00002e030
len=6 cap=10 [0 0 0 0 0 0]
Memory address for v: 0xc000010018
Memory address for v[0]: 0xc0000220a0
```

As seen in the example, the value of `v` does not change after the capacity overflow, because it refers to the **slice variable**.

However, the memory address of the slice elements changes after the capacity overflow, because the slice elements are reallocated to a new memory location.

## Inheriting capacity when slicing a slice


```go
package main

import "fmt"

func main() {
    a := []int{1,2,3}
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("a: %v", a), cap(a), len(a))
    b := a[:1]
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("b: %v", b), cap(b), len(b))
    c := b[:2]
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("c: %v", c), cap(c), len(c))
}
```
```
a: [1 2 3] (cap:3; len:3)
b: [1]     (cap:3; len:1)
c: [1 2]   (cap:3; len:2)
```

In this example, the capacity of the slice `b` is inherited from the slice `a`. Similarly, the capacity of the slice `c` is inherited from the slice `b`. This is because the slice `b` and `c` are created by slicing the slice `a` and `b` respectively.

When doing `c := b[:2]`, the original slice `a` is implicitly referenced by the slice `b`. Therefore, the value `2` will be the second element of `c`.

This can be solved by using the following notation:
```go
package main

import "fmt"

func main() {
    a := []int{1,2,3}
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("a: %v", a), cap(a), len(a))
    b := a[:1]
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("b: %v", b), cap(b), len(b))
    c := b[:2]
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("c: %v", c), cap(c), len(c))
    d := c[0:1:1]
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("d: %v", d), cap(d), len(d))
}
```
```
a: [1 2 3] (cap:3; len:3)
b: [1]     (cap:3; len:1)
c: [1 2]   (cap:3; len:2)
d: [1]     (cap:1; len:1)
```
Here, `d` is created by slicing `c` with a capacity of 1. This way, the capacity of `d` is not inherited from `c`, but is explicitly defined as 1. This is done by using the notation `c[0:1:1]`. The first `1` is the length of the slice, and the second `1` is the capacity of the slice.

Note that every element from a on refers to elements in the original slice. Therefore, if the original slice is modified, the slices created from it will also be modified as in the following example:

```go
package main

import "fmt"

func main() {
    a := []int{1,2,3}
    b := a[:1]
    c := b[:2]
    d := c[0:1:1]
    
    a[0] = 10
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("a: %v", a), cap(a), len(a))
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("b: %v", b), cap(b), len(b))
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("c: %v", c), cap(c), len(c))
    fmt.Printf("%-10s (cap:%d; len:%d)\n", fmt.Sprintf("d: %v", d), cap(d), len(d))
}
```
```
a: [10 2 3] (cap:3; len:3)
b: [10]    (cap:3; len:1)
c: [10 2]  (cap:3; len:2)
d: [10]    (cap:1; len:1)
```