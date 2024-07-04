[Go Class: 06 Control Statements; Declarations & Types](https://www.youtube.com/watch?v=qpHLhmoV3BY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=7)
- [Control Statements](#control-statements)
  - [If-Else](#if-else)
  - [For Loops](#for-loops)
    - [Conditional Loop](#conditional-loop)
    - [Slice/Array range loop](#slicearray-range-loop)
    - [Map range loop](#map-range-loop)
    - [Infinite Loop](#infinite-loop)
    - [Labelled Loop](#labelled-loop)
  - [Switch](#switch)
- [Declarations](#declarations)
  - [Package scope](#package-scope)
  - [Exports](#exports)
  - [Imports](#imports)
  - [Initialization](#initialization)

## Control Statements

### If-Else

```go
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}
```

With short declaration

```go

package main

import "fmt"

func main() {
    if x := 10; x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}
```



### For Loops

#### Conditional Loop
```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```
```
0
1
2
3
4
```

In this example:
- `i := 0` is the initialization statement
- `i < 5` is the check condition,
- `i++` is the post statement.

The loop will run as long as the check condition is true, so the loop will run 5 times.

#### Slice/Array range loop
```go

package main

import "fmt"

func main() {
    a := []int{1, 2, 3}
    // Range loop with index only
    for i := range a {
        fmt.Println(i)
    }
    // Range loop with index and value
    for i, v := range a {
        fmt.Println(i, v)
    }
}
```
#### Map range loop
```go
package main

import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    // Range loop with key only
    for k := range m {
        fmt.Println(k)
    }
    // Range loop with key and value
    for k, v := range m {
        fmt.Println(k, v)
    }
    // Range loop only values
    for _, v := range m {
        fmt.Println(v)
    }
}
```

#### Infinite Loop
```go

package main

import "fmt"

func main() {
    for {
        fmt.Println("Infinite loop")
        if true {
            break
        }
    }
}
```


#### Labelled Loop
```go
package main

import "fmt"

func main() {
    outer:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Println(i, j)
            if i == 2 && j == 2 {
                break outer
            }
        }
    }
}
```


### Switch

Made for selecting between multiple options.

```go

package main

import "fmt"

func main() {

    switch x := 10; x {
    case 1:
        fmt.Println("x is 1")
    case 2:
        fmt.Println("x is 2")
    case 3:
        fmt.Println("x is 3")
    case 4, 5, 6:
        fmt.Println("x is 4, 5, or 6")
    case x > 6:
        fmt.Println("x is greater than 6")
    default:
        fmt.Println("x is not 1, 2, 3, 4, 5, 6, or greater than 6")
    }
}
```

## Declarations

- Every standalone program must have a `main` package.
- Nothing is global in Go, everything is in a package.
- There are 2 scopes in Go: package scope and block scope.

### Package scope
- Anything starting with a keyword can be declared at the package level.
- So a variable declared with `:=` cannot be declared at the package level.

### Exports
- Any name starting with a capital letter is exported, and can be accessed from other packages.
### Imports
- Each source file in a package can only import what it needs.
- Generally, files of the same package are in the same directory.
- There are no circular dependencies in Go, so you can't have a package that imports itself through another package.

### Initialization
- Items within a package get initialized in the order they are declared, before `main` is called.
- You can have a function called `init` in a package, which will be called before `main`. Only the runtime can call `init`, you can't call it yourself.


