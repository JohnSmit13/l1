package main

import (
    "fmt"
    "math"
)

type Point struct {
    x,y float32
}

func main() {
    var x,y Point
    x.MakePoint(11.11, 22.22)
    y = makePoint(33.33, 44.44)
    fmt.Printf("Abstract distance is: %f\n", x.Distance(y))
}

func makePoint(x,y float32) Point {
    return Point{x,y}
}

func(p *Point) MakePoint(x,y float32) {
    p.x = x
    p.y = y
}

func(x *Point) Distance(y Point) float32 {
    var res float32
    var len1, len2 float32
    len1 = (y.x-x.x) * (y.x-x.x)
    len2 = (y.y-x.y) * (y.y-x.y)
    res = float32(math.Sqrt(float64(len1 + len2)))
    return res
}
