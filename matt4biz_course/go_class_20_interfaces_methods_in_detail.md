# [Go Class: 20 Interfaces & Methods in Detail](https://www.youtube.com/watch?v=AXCIEiebVfI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=21)

- [Go Class: 20 Interfaces \& Methods in Detail](#go-class-20-interfaces--methods-in-detail)
  - [An interface variable is `nil` until initialized](#an-interface-variable-is-nil-until-initialized)
    - [Error is really an interface](#error-is-really-an-interface)
  - [Pointer vs value receivers](#pointer-vs-value-receivers)
    - [Consistency on receiver types](#consistency-on-receiver-types)
  - [Currying functions](#currying-functions)
  - [Method values](#method-values)
  - [Interfaces in practice](#interfaces-in-practice)
  - [Empty interfaces](#empty-interfaces)

## An interface variable is `nil` until initialized

An interface has two parts:

- A type
- A pointer to an object of that type

One must be careful when assigning a `nil` value to an interface, as it might not the behavior you expect.

```go
package main

import (
    "bytes"
    "fmt"
    "io"
)

func main() {
    var r io.Reader  // Here r is nil, because it has no type or object
    var b *bytes.Buffer  // Here b is nil, because it has no object
    r=b  // Now r is not nil, because it got a type, even if it has no object. It got initialized.

    fmt.Println(r)  // <nil>
    fmt.Println(r == nil)  // false
}
```

### Error is really an interface

```go
type error interface {
    Error() string
}
```

Error is an interface so that you can define your own error types. The `Error()` method returns a string that describes the error.

```go
package main

import (
    "fmt"
)

type errFoo struct {
    path string
    err error
}

func (e *errFoo) Error() string {
    return fmt.Sprintf("Error in %s: %s", e.path, e.err)
}
```

When using the code above, you can define a function that returns a pointer to `errFoo`. Bút that would be a mistake, as it would return a `nil` pointer to `errFoo`, which is not `nil` because it has a type. So the error check would not work as expected.

```go
func XYZ(i int) *errFoo {
    return nil
}

func main() {
    var err error := XYZ(10)  // BAD, the interface gets a nil concrete pointer.
    if err != nil {
        fmt.Println(err)
    }
}
```

Instead of returning a `nil` pointer, you should return a `nil` interface.

```go

func XYZ(i int) error {
    return nil
}

func main() {
    var err errFoo := XYZ(10)
    if err != nil {
        fmt.Println(err)
    }
}
```

## Pointer vs value receivers

- A method name either has a pointer receiver or a value receiver, but not both.
- Pointer methods may be called on non-pointer values. Go will automatically convert the value to a pointer.

```go

p1 := new(Point)
p2 := Point{1, 2}

// OffsetOf uses a value receiver
p1.Offset(p2)  // same as (*p1).OffsetOf(p2)
p2.Add(3, 4)  // same as (&p2).Add(3, 4)
Point{1, 2}.Add(3, 4)  // Not OK, because Add is a pointer method and Point{1, 2} is a literal.

```

### Consistency on receiver types

- If you have a method with a pointer receiver, all methods should have a pointer receiver. (With some exceptions)
- Objects of that type are generally not safe to be copied.

## Currying functions

Currying takes a function that takes multiple arguments and returns a function that takes fewer arguments.
By doing so, one of the arguments is fixed, and the function can be reused with different values for the other arguments.

```go

func Add(a, b int) int {
    return a + b
}

func AddToA(a int) func(int) int {
    return func(b int) int {
        return Add(a, b)
    }
}

func main() {
    add5 := AddToA(5)
    fmt.Println(add5(3))  // 8
}
```

## Method values

A method value is a function value that has a receiver bound to it.
A method value with a value receiver copies the receiver.
A method value with a pointer receiver copies a pointer to the receiver.

```go
package main

import (
    "fmt"
)

type Point struct {
    X, Y int
}

func (p Point) Add(a, b int) int {
    return p.X + p.Y + a + b
}

func main() {
    p := Point{1, 2}
    add := p.Add  // Method value
    fmt.Println(add(3, 4))  // 10
    p = Point{0, 0}
    fmt.Println(add(3, 4))  // 10, as the method value is closed over the receiver, and the receiver is not changed.
}
```

## Interfaces in practice

1. Let users define interfaces. _What minimal behavior to they require?_
2. Re-use standard interfaces when possible. _io.Reader, io.Writer, etc._
3. Keep interface declarations small. _Don't put everything in one interface._
4. Compose one-method interfaces into larger interfaces. _io.ReaderAt, io.WriterAt, etc._
5. Avoiud coupling interfaces to particular implementations.
6. Accept interfaces, but return concrete types. *Be liberal on what you accept, but conservative on what you return.*


## Empty interfaces

An empty interface is an interface with no methods. It can hold values of any type.

They are often used in functions that can take any type of argument.

```go

func Print(v interface{}) {
    fmt.Println(v)
}

func main() {
    Print(42)
    Print("Hello")
    Print([]int{1, 2, 3})
}
```
