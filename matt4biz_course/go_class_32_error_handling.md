# [Go Class: 32 Error Handling](https://www.youtube.com/watch?v=oIxXp0OgK_0)

- [Go Class: 32 Error Handling](#go-class-32-error-handling)
  - [Errors in Go](#errors-in-go)
  - [Wrapping Errors](#wrapping-errors)
    - [Error.Is()](#erroris)
    - [Error.As()](#erroras)



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

