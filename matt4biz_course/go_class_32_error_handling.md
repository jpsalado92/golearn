# [Go Class: 32 Error Handling](https://www.youtube.com/watch?v=oIxXp0OgK_0)

- [Go Class: 32 Error Handling](#go-class-32-error-handling)
  - [Errors in Go](#errors-in-go)
  - [Wrapping Errors](#wrapping-errors)
    - [Error.Is()](#erroris)
    - [Error.As()](#erroras)
  - [Normal errors](#normal-errors)
  - [Abnormal errors](#abnormal-errors)
  - [Exception handling](#exception-handling)

## Errors in Go

Errors in Go are objects satisfying the error interface.

The error interface is a single method interface with the signature:

```go
type error interface {
  Error() string
}
```

Any concrete type with `Error() string` method satisfies the error interface.

```go
type MyType struct {}

func (e *MyType) Error() string {
  return "MyType error"
}
```

## Wrapping Errors

Custom errors may now unwrap their internal errors.

In order to do so, the custom error must implement the `Unwrap()` method.

```go
type MyType struct {
  err error
}

func (e *MyType) Error() string {
  return e.err
}
```

### Error.Is()

`Error.Is` is used to check whether an error has another error **variable** in its chain.

```go
if audio, err = DecodeWaveFile("test.wav"); err != nil {
  if errors.Is(err, os.ErrPermission) {
    // Report security breach
    ...
  }
  ...
}
```

So, if we export our error variables, we can implement the `Is()` method in custom
variable types to check for specific errors.

```go
type WaveError struct {
  kind errKind
  ...
}

func (w *WaveError) Is(t, error) bool {
  e, ok := t.(*WaveError)  // Reflection

  if !ok {
    return false
  }

  return w.kind == e.kind
}
```

### Error.As()

`Error.As` is used to check whether an error has another error **type** in its chain.

```go
if audio, err = DecodeWaveFile("test.wav"); err != nil {
  var e os.PathError // a struct type

  if errors.As(err, &e) {
    // Here you might decide to just pass the underlying error to the caller
    return e
    ...
  }
  ...
}
```

## Normal errors

Result from input or external conditions.

In Go these are handled by returning an error object.

```go
// Not exactly os.Open, but shows the basic logic
func Open(name string, flag int, perm FileMode) (*File, error) {
  r, e := syscall.Open(name, flag|syscall.O_CLOEXEC, syscallMode(perm))
  if e != nil {
    return nil, &PathError{"open", name, e}
  }
  return newFile(uintptr(r), name, kindOpenFile), nil
}
```

## Abnormal errors

Result from invalid programming logic. Could be considered as bugs.

```go
func (d *digest) checkSum() [Size]byte {
// finish writing the checksum
  . . .
  if d.nx != 0 { // panic if there's data left over
    panic("d.nx != 0")
  }
. . .
}
```

This should be used when assumptions of our own programming design are wrong.


## Exception handling

Other programming languages like Ada, C++, Java, Python, etc. have exception handling, so that the program can recover from an error.

Go does not support exception handling. However, it can the `panic` and `recover` mechanism to simulate exception handling.

```go
func abc() {
  panic("omg")
}

func main() {
  defer func() {

if p := recover(); p != nil {
    // what else can you do?
    fmt.Println("recover:", p)
  }
  }()

abc()
}
```

In the above code, the `defer` function will be called when the `abc()` function panics.
The `recover()` function will return the value that was passed to the `panic()` function.