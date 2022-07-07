package main

import (
	"fmt"
)

func str() int {
	str := "123でゴール"

  for _, i := range str {
    fmt.Println("rune:", i)
    fmt.Println("rune to String:",string(i))
  }
  i := 1
  return i
}
