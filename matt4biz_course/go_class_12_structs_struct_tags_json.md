# [Go Class: 12 Structs, Struct tags & JSON](https://www.youtube.com/watch?v=0m6iFd9N_CY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=14)

- [Go Class: 12 Structs, Struct tags \& JSON](#go-class-12-structs-struct-tags--json)
  - [Struct declaration](#struct-declaration)
  - [Struct properties](#struct-properties)
    - [Structs are value types](#structs-are-value-types)
    - [Struct comparison](#struct-comparison)
    - [Zero value of a struct](#zero-value-of-a-struct)
    - [Copying structs](#copying-structs)
  - [Defining structs](#defining-structs)
  - [Nested structs](#nested-structs)
    - [Value nested struct](#value-nested-struct)
    - [Nested struct pointers](#nested-struct-pointers)
  - [Anonymous structs](#anonymous-structs)
  - [Struct tags](#struct-tags)
  - [Gotchas](#gotchas)
    - [1. `map[string]struct` vs `map[string]*struct`](#1-mapstringstruct-vs-mapstringstruct)
    - [2. Passing down structs as parameters vs pointers](#2-passing-down-structs-as-parameters-vs-pointers)
    - [3. Using a map of empty struct as set](#3-using-a-map-of-empty-struct-as-set)

## Struct declaration

A `struct` is a composite data type that groups together zero or more named values of arbitrary types as fields. Each field in a struct has a name and a type. The fields are declared in a comma-separated list within curly braces `{}`. The fields can be accessed using the `.` operator.

```go

package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {
    p := person{name: "Alice", age: 25}
    fmt.Println(p)
    fmt.Println(p.name)
    fmt.Println(p.age)
}
```

## Struct properties

### Structs are value types

This means that when they are assigned to a new variable or passed as a parameter to a function, a copy of the struct is made. This is different from reference types, such as slices and maps, where changes to the value are reflected in all references to the value.

### Struct comparison

You can only compare structs if they share the same type.

Similar struct types might be compared if you cast them to the same type, but in order for that to work, the fields must be the same and in the same order.

### Zero value of a struct

The zero value of a struct is a struct with all its fields set to their zero values.

For example, the zero value of a struct with a string field is a struct with an empty string field, and the zero value of a struct with an int field is a struct with an int field set to 0.

### Copying structs

You can copy a struct by assigning it to a new variable of the same type. This creates a new copy of the struct with the same field values.

## Defining structs

```go
package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {
    // Literal whole
    p1 := person{"Alice", 25}
    fmt.Println(p1)

    // Literal partial
    p2 := person{name: "Bob"}
    fmt.Println(p2)

    // Field by field
    p3 := person{}
    p3.name = "Charlie"
    p3.age = 30
    fmt.Println(p3)
}
```

```
{Alice 25}
{Bob 0}
{Charlie 30}
```

## Nested structs

### Value nested struct

In this example, the `person` struct has a field `address` that is an `address` struct. This means that the `address` field in the `person` struct will store a copy of the `address` struct.

```go
package main

import "fmt"

type address struct {
    city  string
    state string
}

type person struct {
    name    string
    age     int
    address address
}

func main() {
    address := address{
        city:  "New York",
        state: "NY",
    }
    p := person{
        name: "Alice",
        age:  25,
        address: address,
    }
    fmt.Println(p)
    fmt.Println(p.address.city)
    fmt.Println(p.address.state)
}
```

### Nested struct pointers

In this example, the `person` struct has a field `address` that is a pointer to an `address` struct. This means that the `address` field in the `person` struct will store the memory address of an `address` struct.

```go

package main

import "fmt"

type address struct {
    city  string
    state string
}

type person struct {
    name    string
    age     int
    address *address
}

func main() {
    address := address{
        city:  "New York",
        state: "NY",
    }
    p := person{
        name: "Alice",
        age:  25,
        address: &address,
    }
    fmt.Println(p)
    fmt.Println(p.address.city)
    fmt.Println(p.address.state)
}
```

## Anonymous structs

```go

package main

import "fmt"

func main() {
    p := struct {
        name string
        age  int
    }{
        name: "Alice",
        age:  25,
    }
    fmt.Println(p)
    fmt.Println(p.name)
    fmt.Println(p.age)
}
```

You can compare anonymous structs if they have the same fields and types in the same order.

```go
package main

import "fmt"

func main() {
    p1 := struct {
        name string
        age  int
    }{
        name: "Alice",
        age:  25,
    }
    p2 := struct {
        name string
        age  int
    }{
        name: "Jon",
        age:  5,
    }
    fmt.Println(p1 == p2)
}
```

## Struct tags

- Struct tags are key-value pairs that are associated with struct fields.

## Gotchas

### 1. `map[string]struct` vs `map[string]*struct`

It is preferable to use `map[string]*struct` instead of `map[string]struct` when you want to store pointers to structs in a map. This is because when you store a struct in a map, a copy of the struct is made, which can be inefficient if the struct is large. Storing a pointer to the struct in the map avoids this overhead.

When structs are stored in a map by using a `map[string]struct` map, it is unsafe to pick the pointer of any of the struct values of the map. This is due to the fact that the map can be rehashed, which can cause the pointer to become invalid.


### 2. Passing down structs as parameters vs pointers

```go

package main

import "fmt"

type person struct {
  name string
  age  int
}

func modifyPerson(p person) {
  p.name = "Bob"
  p.age = 30
}

func modifyPersonPointer(p *person) {
  p.name = "Charlie"
  p.age = 35
}

func main() {
  p := person{name: "Alice", age: 25}
  fmt.Println(p) // {Alice 25}
  modifyPerson(p)
  fmt.Println(p) // {Alice 25}
  modifyPersonPointer(&p)
  fmt.Println(p) // {Charlie 35}
}
```

### 3. Using a map of empty struct as set

```go

package main

import "fmt"

func main() {
  s := make(map[string]struct{})
  s["Alice"] = struct{}{}
  s["Bob"] = struct{}{}
  fmt.Println(s)
  fmt.Println("Alice" in s)
  fmt.Println("Charlie" in s)
}
```

```go

package main

import "fmt"

func main() {
  s := make(map[string]struct{})
  s["Alice"] = struct{}{}
  s["Bob"] = struct{}{}
  fmt.Println(s)
  fmt.Println("Alice" in s)
  fmt.Println("Charlie" in s)
}
```
