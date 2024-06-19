# [Go Class: 03 Basic Types](https://www.youtube.com/watch?v=NNLpEPb2ddE&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=4)
- [Go Class: 03 Basic Types](#go-class-03-basic-types)
  - [Basic Types](#basic-types)
    - [Integers](#integers)
    - [Real and Complex Numbers](#real-and-complex-numbers)
    - [Text Types](#text-types)
    - [Special types](#special-types)
  - [Multiple ways to declare variables](#multiple-ways-to-declare-variables)
  - [Showing the type and value of something](#showing-the-type-and-value-of-something)
  - [Variables initialization](#variables-initialization)
  - [Constants](#constants)

## Basic Types
### Integers
Unsized integers `int` default to the machine's word size (32 or 64 bits).

`int` is the default type for integers in go.

There are many types of integers in Go:
**Signed**: `int8`, `int16`, `int32`, `int64`
**Unsigned**: `uint8`, `uint16`, `uint32`, `uint64`

(Unsigned means that the number is always positive.)


### Real and Complex Numbers

**Real numbers**: `float32`, `float64`

**Complex numbers**: `complex64`, `complex128`

For monetary calculations better use the [Go money](https://pkg.go.dev/github.com/Rhymond/go-money) package.

### Text Types
### Special types

`bool` (true or false), these are not convertible from integers an in some other languages.

`error` A special type with a function `Error()` that returns a string. Its value might be `nil` or `non-nil`.

Pointers are a special type in Go. They are used to store the memory address of a value. They might be `nil` or `non-nil`. There is no way to manipulate memory addresses in Go unless you use the `unsafe` package.







## Multiple ways to declare variables
    
```go
var a int

var b int = 10

c := 20  // Only works inside a function

var (
    d int
    e int = 30
    f = 40
)
```


In go `:= ` is the short declaration operator. It is used to declare and initialize a variable.


## Showing the type and value of something

```go
fmt.Printf("%T %v\n", a, a)
```

## Variables initialization
There are no non-initialized variables in Go. If you declare a variable and don't assign a value to it, it will be initialized with the zero value of its type.

- All numerical values get initialized to 0.
- `bool` gets initialized to `false`.
- `string` gets initialized to `""`.
- Everything else gets initialized to `nil`.
- For aggregate types (arrays, slices, structs, etc.) the zero value is the zero value of all of its elements.

## Constants
Constants are declared with the `const` keyword, and they are immutable.
  

```go


