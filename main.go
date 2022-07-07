package main

import (
	"fmt"
	"math"
)

type Point struct {
  X intAAA
  Y int
}

func main() {  
  countContry()
  fmt.Println("end")
}

func (from *Point) Distance(to Point) float64 {
  sum := math.Pow(float64(to.X - from.X), 2) + math.Pow(float64(to.Y - from.Y), 2) 
  fmt.Println("平方根の中身：", sum)
  return math.Sqrt(sum)
}

