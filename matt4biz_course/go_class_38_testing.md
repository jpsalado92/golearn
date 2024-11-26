# [Go Class: 38 Testing](https://www.youtube.com/watch?v=PIPfNIWVbc8)

## Basic Rules

Test files names end with `_test.go`

Test files contain test functions that are named starting with `Test`

In order to run tests, we use the `go test` command

Tests are cached, so if no changes are done to the code, the cached values will be used.
You can use the `-count=1` flag to run the tests without using the cache.

Every test function has the same signature `func TestXxx(t *testing.T)`. The errors are
reported using the `t.Error` method.

## Table Driven Tests

We can have table driven tests, where we have a slice of structs that contain the input and
expected output values. We then iterate over the slice and run the tests.

```go
package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {3, 4, 7},
        {5, 6, 11},
        {7, 8, 15},
    }

    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf("Expected %d, but got %d", test.expected, result)
        }
    }
}
```

## Table Driven Subtests

These are different from table driven tests in the way the errors are reported. In this
case, we use the `t.Run` method to create subtests, which takes a name for the subtest and
a function

```go

func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {3, 4, 7},
        {5, 6, 11},
        {7, 8, 15},
    }

    for _, test := range tests {
        t.Run(fmt.Sprintf("%d+%d", test.a, test.b), func(t *testing.T) {
            result := Add(test.a, test.b)
            if result != test.expected {
                t.Errorf("Expected %d, but got %d", test.expected, result)
            }
        })
    }
}
```

## Main testing functions

These are root functions that are run before and after the tests. They are used to setup
and teardown resources needed for the tests.

```go

func TestMain(m *testing.M) {
    stop, err := setup()
    if err != nil {
        log.Fatal(err)
    }
    result := m.Run()
    stop()
    os.Exit(result)
}
```

## Special test only packages

We can have test only packages that are used to test the main package. These packages are
named with the `_test` suffix. They can be used to test unexported functions in the main
package.

These packages are not included in the final binary.

https://pkg.go.dev/testing#pkg-types:~:text=The%20test%20file%20can%20be%20in%20the%20same%20package%20as%20the%20one%20being%20tested%2C%20or%20in%20a%20corresponding%20package%20with%20the%20suffix%20%22_test%22.