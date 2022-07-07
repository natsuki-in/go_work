package main

import (
	"fmt"
	"math"
)

type Point struct {
  X int
  Y int
}
]
func main() {  
  p1 := Point{X: 8, Y:16,}
  p2 := Point{X: 3, Y: 4,}
  fmt.Println("距離：", p1.Distance(p2))
}

func (from *Point) Distance(to Point) float64 {
  sum := math.Pow(float64(to.X - from.X), 2) + math.Pow(float64(to.Y - from.Y), 2) 
  fmt.Println("平方根の中身：", sum)
  return math.Sqrt(sum)
}

