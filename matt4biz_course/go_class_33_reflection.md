# [Go Class: 33 Reflection](https://www.youtube.com/watch?v=T2fqLam1iuk)

- [Go Class: 33 Reflection](#go-class-33-reflection)
  - [What does reflection mean](#what-does-reflection-mean)
  - [Empty interface](#empty-interface)
  - [Type assertion, downcasting](#type-assertion-downcasting)
  - [Syntax](#syntax)
  - [Reflection package](#reflection-package)
  - [Type switch](#type-switch)
  - [Example 1 custom JSON decoder](#example-1-custom-json-decoder)
  - [Example 2 Check values within a JSON piece while unmarshalling](#example-2-check-values-within-a-json-piece-while-unmarshalling)

## What does reflection mean

Reflection is the ability of a program to examine its own structure, particularly types.

## Empty interface

The empty interface `interface{}` says nothing

- It does not not define any methods
- It does not define any behavior
- It does not restrict any behavior

So it can be used to **represent any type** in Go.

## Type assertion, downcasting

Some container types use the empty interface to store values of different types. To get the value back, you need to use a type assertion through reflection.

```go
// Type assertion
func main() {
  c := Container{Value: 42}
  i, ok := c.Value.(int)
  if !ok {
    fmt.Println("Value is not an int")
  }
  fmt.Println(i)
}
```

## Syntax

To use type assertion, you need to use the following syntax:

```go
// Container type
type Container struct {
  Value interface{}
}
value, ok := container.(Type)
```

Where

- `container` is the variable of the container type
- `Type` is the type you want to assert
- `value` is the variable that will hold the value if the assertion is successful
- `ok` is a boolean that will be `true` if the assertion is successful.

Alternatively, you can avoid using ok, which will cause a `panic` if the assertion fails.

## Reflection package

Deep equality is a common use case for reflection. The `reflect` package provides a way to compare two values of different types.

In this example, a string and a slice of integers are compared. They are normally not comparable with `==`, as they are of different types.

```go
import (
  "fmt"
  "reflect"
)

func main() {
  a := "a string"
  b := []int{1, 2, 3}
  fmt.Println(reflect.DeepEqual(a, b))
}
```

In this case `DeepEqual` function returns `false` because the contents are different.

But for this other example, the function returns `true` because the contents are the same even though the types are different.

```go
import (
  "fmt"
  "reflect"
)

func main() {
  a := []byte("hello")
  b := "hello"
  fmt.Println(reflect.DeepEqual(a, []byte(b))) // true
}
```

## Type switch

We can also use type assertion in a switch statement to check the type of a variable, not its value.

```go
func main() {
  var i interface{} = 42
  switch v := i.(type) {
  case int:
    fmt.Println("int", v)
  case string:
    fmt.Println("string", v)
  default:
    fmt.Println("unknown")
  }
}
```

## Example 1 custom JSON decoder

Not all JSON messages are well-behaved. What if some keys depend on others in the message?

```json
{
"item": "album",
"album": {"title": "Dark Side of the Moon"}
}
{
"item": "song",
"song": {"title": "Bella Donna", "artist": "Stevie Nicks"}
}

```

```go

type response struct {
  Item string `json:"item"`
  Album string
  Title string
  Artist string
}
type respWrapper struct { // We need respWrapper because it must have a separate unmarshal method from the response type (see below)
  response
}

func (r *respWrapper) UnmarshalJSON(b []byte) (err error) {
  var raw map[string]interface{}
  err = json.Unmarshal(b, &r.response) // ignore error handling
  err = json.Unmarshal(b, &raw)
  switch r.Item {
  case "album":
    inner, ok := raw["album"].(map[string]interface{})
    if ok {
      if album, ok := inner["title"].(string); ok {
        r.Album = album
      }
    }
  case "song":
    inner, ok := raw["song"].(map[string]interface{})
    if ok {
      if title, ok := inner["title"].(string); ok {
        r.Title = title
      }
      if artist, ok := inner["artist"].(string); ok {
        r.Artist = artist
      }
    }
  }
  return err
}

var j1 = `{
  "item": "album",
  "album": {"title": "Dark Side of the Moon"}
}`

var j2 = `{
  "item": "song",
  "song": {"title": "Bella Donna", "artist": "Stevie Nicks"}
}`

func main() {
  var resp1, resp2 respWrapper
  var err error

  if err = json.Unmarshal([]byte(j1), &resp1); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%#v\n", resp1.response)

  if err = json.Unmarshal([]byte(j2), &resp2); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%#v\n", resp2.response)
}

// main.response{Item:"album", Album:"Dark Side of the Moon",
// Title:"", Artist:""}
//
// main.response{Item:"song", Album:"", Title:"Bella Donna",
// Artist:"Stevie Nicks"}
```

## Example 2 Check values within a JSON piece while unmarshalling

We want to know if a known fragment of JSON is contained in a larger unknown piece

`{"id": "Z"}` in? `{"id": "Z", "part": "fizgig", "qty": 2}`

```go
func matchNum(key string, exp float64, data map[string]interface{}) bool {
if v, ok := data[key]; ok {
if val, ok := v.(float64); ok && val == exp {
return true
}
}
return false
}

func matchString(key, exp string, data map[string]interface{}) bool {
  // is it in the map?
  if v, ok := data[key]; ok {
    // is it a string, and does it match?
    if val, ok := v.(string); ok && strings.EqualFold(val, exp) {
      return true
    }
  }
  return false
}

func contains(exp, data map[string]interface{}) error {
  for k, v := range exp {

    switch x := v.(type) {

    case float64:
      if !matchNum(k, x, data) {
        return fmt.Errorf("%s unmatched (%d)", k, int(x))
      }

    case string:
      if !matchString(k, x, data) {
        return fmt.Errorf("%s unmatched (%s)", k, x)
      }

    case map[string]interface{}:
      if val, ok := data[k]; !ok {
        return fmt.Errorf("%s missing in data", k)
      } else if unk, ok := val.(map[string]interface{}); ok {
        if err := contains(x, unk); err != nil {
          return fmt.Errorf("%s unmatched (%+v): %s", k, x, err)
        }
      } else {
        return fmt.Errorf("%s wrong in data (%#v)", k, val)
      }
    }
  }
  return

func CheckData(want, got []byte) error {
  var w, g map[string]interface{}

  if err := json.Unmarshal(want, &w); err != nil {
    return err
  }

  if err := json.Unmarshal(got, &g); err != nil {
    return err
  }
  return contains(w, g)
}

```

Write tests for the function

```go
var unknown = `{
  "id": 1,
  "name": "bob",
  "addr": {
    "street": "Lazy Lane",
    "city": "Exit",
    "zip": "99999"
  },
  "extra": 21.1
}`

func TestContains(t *testing.T) {
  var known = []string{
    `{"id": 1}`,
    `{"extra": 21.1}`,
    `{"name": "bob"}`,
    `{"addr": {"street": "Lazy Lane", "city": "Exit"}}`,
  }
  for _, k := range known {
    if err := CheckData(k, []byte(unknown)); err != nil {
      t.Errorf("invalid: %s (%s)\n", k, err)
    }
  }
}

func TestNotContains(t *testing.T) {
  var known = []string{
    `{"id": 2}`,
    `{"pid": 2}`,
    `{"name": "bobby"}`,
    `{"first": "bob"}`,
    `{"addr": {"street": "Lazy Lane", "city": "Alpha"}}`,
  }
  for _, k := range known {
    if err := CheckData(k, []byte(unknown)); err == nil {
      t.Errorf("false positive: %s\n", k)
    } else {
      t.Log(err)
    }
  }
}
```

Test the function with

```bash
go test -v
go test ./... -cover
go test ./... -coverprofile=c.out -covermode=count
go tool cover -html=c.out
```
