# (Go Class: 19 Composition)[https://www.youtube.com/watch?v=0X6AcnwocbM&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=20]

- [(Go Class: 19 Composition)\[https://www.youtube.com/watch?v=0X6AcnwocbM\&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6\&index=20\]](#go-class-19-compositionhttpswwwyoutubecomwatchv0x6acnwocbmlistploilbko9rg3skrcj37kn5zj803hhiurk6index20)
  - [Hints](#hints)
  - [Field promotion](#field-promotion)
  - [Example 1: Using composition to sort a list of tuples](#example-1-using-composition-to-sort-a-list-of-tuples)
  - [Good practice 1: Make nil useful](#good-practice-1-make-nil-useful)

## Hints

- Try to get interfaces as simple as possible. The smaller methods the better.
- Use composition to create complex types from simple ones.
- Composition is not inheritance.

## Field promotion

The attributes and methods of an embedded struct are promoted to the outer struct, so the outer struct can access them directly.

If fields are redefined in the outer struct, the outer struct's references will be used (they get promoted), but the embedded struct's fields can still be accessed with the embedded struct's name.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func (p *Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

type Employee struct {
    Person
    Title string
    Name string
}

func (e *Employee) Greet() {
    fmt.Println("Hello Sir, my name is", e.Name)
}

func main() {
    e := Employee{
        Person: Person{
            Name: "John",
        },
        Title: "Manager",
        Name: "Employee",
    }
    e.Age = 30
    fmt.Println(e.Name)
    fmt.Println(e.Person.Name)
    e.Greet()  // Hello Sir, my name is Employee
    e.Person.Greet()  // Hello, my name is John
}
```

Promotion also works when a struct embeds a pointer to another struct.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func (p *Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

type Employee struct {
    *Person
    Title string
}

func main () {
    e := Employee{
        Person: &Person{
            Name: "John",
        },
        Title: "Manager",
    }
    e.Age = 30
    fmt.Println(e.Name)
    fmt.Println(e.Person.Name)
    e.Greet()  // Hello, my name is John
}
```

## Example 1: Using composition to sort a list of tuples

embedding an interface to a struct

```go

package main

import (
    "fmt"
    "sort"
)

type Organ struct {
    name string
    weight int
}

type Organs []Organ

func (o Organs) Len() int {
    return len(o)
}

func (o Organs) Swap(i, j int) {
    o[i], o[j] = o[j], o[i]
}

type ByWeight struct {
    Organs
}

func (b ByWeight) Less(i, j int) bool {
    return b.Organs[i].weight < b.Organs[j].weight
}

type ByName struct {
    Organs
}

func (b ByName) Less(i, j int) bool {
    return b.Organs[i].name < b.Organs[j].name
}

func main() {
    organs := Organs{
        {"heart", 300},
        {"liver", 200},
        {"kidney", 150},
    }
    fmt.Println(organs)
    sort.Sort(ByWeight{organs})
    fmt.Println(organs)
    sort.Sort(ByName{organs})
    fmt.Println(organs)
}
```

## Good practice 1: Make nil useful

Nothing in Go prevents calling a method with a `nil` receiver.

```go
type IntList struct {
    Value int
    Tail *IntList
}

// Sum returns the sum of all elements in the list
func (il *IntList) Sum() int {
    if il == nil {
        return 0 // This is a good practice
    }
    return il.Value + il.Tail.Sum()
}

func main() {
    var il *IntList
    fmt.Println(il.Sum())  // 0
}
```
