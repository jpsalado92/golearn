# [Go Class: 31 Odds & Ends](https://www.youtube.com/watch?v=oTtYtrFv3gw&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=31)

- [Go Class: 31 Odds \& Ends](#go-class-31-odds--ends)
  - [Enumerated Types](#enumerated-types)
  - [Variable Argument Lists](#variable-argument-lists)
  - [Unsigned Integers](#unsigned-integers)
  - [Shortening integers](#shortening-integers)
  - [Goto](#goto)

Topics: #enumeratedTypes #variableArgumentLists #unsignedIntegers #bitWiseOperators

## Enumerated Types

There are no enumerated types in Go. Instead, we use named types and constants.

```go
package main

import "fmt"

type color int

const (
    red color = iota
    blue
    green
)

func main() {
    c := red
    fmt.Println(c)
}
```

There are smart ways to use iota to create a sequence of values. For example, we can use it to create a sequence of powers of 2.

```go

package main

import "fmt"

type bit int

const (
    _ bit = 1 << (10 * iota)
    kb
    mb
    gb
    tb
)

func main() {
    fmt.Println(kb, mb, gb, tb)
}
```

## Variable Argument Lists

These are used when you don't know how many arguments you will be passing to a function.
Note that only the last parameter of a function can be a variable argument list.

```go

package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum())
    fmt.Println(sum(1))
    fmt.Println(sum(1, 2, 3, 4, 5))
    // You can pass a slice to a function that takes a variable argument list by using the `...` operator.
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(sum(nums...))
    // You can leverage this behavior and use append to add elements to a slice.
    nums = append(nums, nums...  )  // This will double the elements in the slice.
    nums = append(nums, 6, 7, 8, 9, 10)  // This will add 5 more elements to the slice.
    fmt.Println(sum(nums...))

}

```

## Unsigned Integers

Sometimes you are required to use low-level protocols (TCP/IP, etc.).
So it is necessary to understand how to work with unsigned integers.

```go

type TCPFields struct {
    SourcePort      uint16
    DestPort        uint16
    SeqNum          uint32
    AckNum          uint32
    DataOffset      uint8
    Reserved        uint8
    Flags           uint8
    WindowSize      uint16
    Checksum        uint16
    UrgentPointer   uint16
}

```

## Shortening integers

When converting integers from uint32 to uint16 in Go, the first 16 bits are retained and the rest are discarded.
This could lead to unexpected results if you are not careful.

## Goto

The `goto` statement is rarely used in Go. It is used to jump to a label in the code.

Every once in a long while, goto is simply easier to understand

```go
readFormat:
    err = binary.Read(buf, binary.BigEndian, &header.format)
    if err != nil {
        return &header, nil, HeaderReadFailed.from(pos, err)
    }
    if header.format == junkID {
        . . . // find size & consume WAVE junk header
        goto readFormat
    }
    if header.format != fmtID {
        return &header, nil, InvalidChunkType
    }
```
