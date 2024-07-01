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