# [Go Class: 18 Methods and Interfaces](https://www.youtube.com/watch?v=W3ZWbhQF6wg&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=20)

- [Go Class: 18 Methods and Interfaces](#go-class-18-methods-and-interfaces)
  - [What is a Method?](#what-is-a-method)
  - [What is an Interface?](#what-is-an-interface)
  - [Why interfaces?](#why-interfaces)
  - [Interfaces and Structural Typing](#interfaces-and-structural-typing)
  - [Properties of methods](#properties-of-methods)
  - [Example 1: Geometry (Point, Line, Path) printing](#example-1-geometry-point-line-path-printing)
- [Example 2: Colored points](#example-2-colored-points)

## What is a Method?

**Methods are type-bound functions**: A special type of a function that has a receiver. The receiver is a type that the method is attached to, and it is placed before the function name.

```go
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

## What is an Interface?

**A specification that lists a set of methods**. Types can satisfy an interface by implementing the methods in the interface.

```go
type Stringer interface {
    String() string
}

func (p Point) String() string {
    return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}  // Point satisfies the Stringer interface
```

## Why interfaces?

- **Code reuse**: You can specify functions in an interface and then reuse them in multiple types through the interface.
- **Abstraction**; Can come up with abstractions that fit many types.
- **Polymorphism**: You can write functions that take an interface as an argument, and then you can pass in any type that satisfies that interface.
- **Decoupling**: You can decouple the implementation of a type from the functions that operate on it.
- **Testing**: You can write tests that use interfaces to test multiple types.
- **Documentation**: You can document the behavior of a type by documenting the interface it satisfies.
- **Extensibility**: You can define new types that satisfy an existing interface.

## Interfaces and Structural Typing

An interface specifies required behavior as a **method set**. Any type that implements that **method set** satisfies the interface. This is known as structural typing, or duck typing. _If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck._

So:

- You don't need to explicitly say that a type implements an interface.
- You can define an interface that a type implements without changing the type.

## Properties of methods

1. **All methods of a given type must be declared in the same package** where the type is declared. You can always extend a given type with new methods by defining a new type that embeds the original type, so you can actually use that to define more methods in another package.

2. **Methods are definable on any user-declared (named) type**. Not only structs (There are some exceptions see package insert for details)

```go
// A method cannot be declared on int, but it can be declared on a named type.
type MyInt int

func (mi MyInt) String() string {
    return fmt.Sprintf("MyInt: %d", mi)
}

// A method can be declared on a struct
type Point struct {
    X, Y int
}

func  (p Point) String() string {
    return fmt.Sprintf("Point: (%d, %d)", p.X, p.Y)
}
```

3. Methods can **take either a value or a pointer receiver**, but not both

```go
func (ml *MyList) Add(n int) { // This is a pointer receiver
    *ml = append(*ml, n)
}

func (ml MyList) Count() int { // This is a value receiver
    return len(ml)
}
```

## Example 1: Geometry (Point, Line, Path) printing

```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	P1, P2 Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.P2.X-l.P1.X, l.P2.Y-l.P1.Y)
}
func (l *Line) ScaleBy(f float64) {
  l.P2.X = l.P1.X + (l.P2.X-l.P1.X)*f
  l.P2.Y = l.P1.Y + (l.P2.Y-l.P1.Y)*f
}

type Path []Point

func (p Path) Distance() (sum float64) {
	sum = 0.0
	for i := 0; i < len(p)-1; i++ {
		sum += Line{p[i], p[i+1]}.Distance()
	}
	return sum
}

func PrintDistance(d Distancer) {
	fmt.Println(d.Distance())
}

type Distancer interface {
	Distance() float64
}

func main() {
	side := Line{Point{1, 2}, Point{4, 6}}
	PrintDistance(side)
	side.ScaleBy(2)
	PrintDistance(side)

	path := Path{
		{1, 1}, // No need to specify Point, as it is inferred from the type of Path
		{5, 1},
		{5, 4},
		{1, 1},
	}
	PrintDistance(path)
}
```

# Example 2: Colored points

```go
package main

import (
  "fmt"
  "image/color"
  "math"
)

type Point struct {
  X, Y float64
}

func (p Point) Distance(q Point) float64 {
  return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
  Point
  Color color.RGBA
}

func main() {
  var cp ColoredPoint
  cp.X = 1
  fmt.Println(cp.Point.X) // 1
  cp.Point.Y = 2
  fmt.Println(cp.Y) // 2
  fmt.Println(cp.Point.Distance(Point{1, 2})) // 0
}
```
