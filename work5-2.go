package main

import (
	"fmt"
	"math"
)

func circle() float64{
  ratio := 3.14
  long := 5.0
  half_long := long / 2

  fmt.Println(half_long)
  m2 := math.Pow(half_long, 2) * ratio

  return m2
}


