# [Go Class: 17 Go does OOP](https://www.youtube.com/watch?v=jexEpE7Yv2A&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=18)

- [Go Class: 17 Go does OOP](#go-class-17-go-does-oop)
  - [OOP properties](#oop-properties)
    - [Abstraction](#abstraction)
    - [Encapsulation](#encapsulation)
    - [Polymorphism](#polymorphism)
      - [Protocol-oriented polymorphism example](#protocol-oriented-polymorphism-example)
    - [Inheritance](#inheritance)
      - [Example](#example)
  - [OOP in Go](#oop-in-go)
  - [Classes in Go](#classes-in-go)

## OOP properties

### Abstraction

Decoupling behavior from implementation details.

As an example, the UNIX file system is an abstraction, roughly 5 different functions to read and write files, but the implementation details are hidden:

- `open`
- `close`
- `read`
- `write`
- `ioctl`

### Encapsulation

Hiding implementation details from misuse.

It's hard to maintain abstraction if the details are exposed:

- The internals may be manipulated in ways contrary to the design.
- Users may become dependent on the details, but those might change.

Encapsulation usually means controlling the visibility of names ("private" variables and methods).

### Polymorphism

`Poly = many, morph = form.`; Multiple types behind a single interface.

Three main types of polymorphism:

- **Ad-hoc polymorphism**: Overloading functions. Meaning the same function name can be used with different types of arguments.
- **Parametric polymorphism**: Generics. Meaning the same function can be used with different types of arguments.
- **Subtype polymorphism**: Subclasses can be used where the superclass is expected.

There is something called "protocol-oriented polymorphism" in Go, which is a way to achieve polymorphism without inheritance. It's a way to define a set of methods that a type must implement to be considered a member of a certain interface.

#### Protocol-oriented polymorphism example

```go
type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{
        Dog{"Fido"},
        Cat{"Whiskers"},
    }

    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

In this example, `Dog` and `Cat` are not related in any way, but they both implement the `Animal` interface, which requires a `Speak` method. This allows them to be treated as `Animal`s in the `main` function.

### Inheritance

Inheritance is a way to reuse code and create a hierarchy of types.

There is a big discussion about how inheritance is not a good idea because it creates a tight coupling between the parent and child classes. In Go, composition is used to achieve the same goal without the tight coupling. **Not having inheritance means better encapsulation and isolation.**

See:

- [Composition over inheritance](https://en.wikipedia.org/wiki/Composition_over_inheritance)
- The inheritance tax (The Pragmatic Programmer)

Some terms:
**Duck typing**: If it looks like a duck and quacks like a duck, it's a duck.
**Composition**: Embedding a type in another type.

#### Example

In the following example, the `Polygon` interface defines a method `Area() float64`. The `Circle`, `Rectangle`, and `Triangle` types all implement the `Polygon` interface, so they can be treated as `Polygon`s in the `main` function.

Composition is achieved by embedding the `Circle`, `Rectangle`, and `Triangle` types in the `Polygon` interface.
Duck typing is shown by the fact that the `Circle`, `Rectangle`, and `Triangle` types don't have to be related in any way to implement the `Polygon` interface.

```go
type Polygon interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Triangle struct {
    Base   float64
    Height float64
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

func main() {
    polygons := []Polygon{
        Circle{Radius: 2},
        Rectangle{Width: 3, Height: 4},
        Triangle{Base: 5, Height: 6},
    }

    for _, polygon := range polygons {
        fmt.Println(polygon.Area())
    }
}
```

## OOP in Go

- Encapsulation is achieved by controlling the visibility of names. Upper-case names are exported, lower-case names are not.
- Abstraction & Polymporphism are achieved through interfaces.
- Inheritance is achieved through composition. Instead of creating a subclass, you embed a type in another type. This is a way to reuse code and create a hierarchy of types without the tight coupling of inheritance.

Go does not offer inheritance.

## Classes in Go

**Go does not have classes.**

Instead, it has types and methods. A method is a function with a receiver argument. The receiver argument is the type that the method is attached to.

```go

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func main() {
    circle := Circle{Radius: 2}
    fmt.Println(circle.Area())
}
```

In this example, `Area` is a method attached to the `Circle` type. The receiver argument `c` is the `Circle` type, so the `Area` method can access the `Radius` field of the `Circle` type.

Go can implement methods in any type, including built-in types like `int` and `string`.

```go
type MyInt int

func (i MyInt) Double() MyInt {
    return i * 2
}

func main() {
    myInt := MyInt(3)
    fmt.Println(myInt.Double())
}
```
