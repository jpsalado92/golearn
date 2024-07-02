package main

type Point struct {
  X, B string
}
type PointA struct {
  Z, W float64
  Point
}
type PointB struct {
  A, B float64
  PointA
}

func main() {
  var cp PointB
  cp.Z = 3
  cp.W = 4
  cp.A = 5
  cp.B = 6
  cp.B = "6"
}