package main

import "fmt"

func check_uruudosi(year int) {
  if ((year % 4 == 0) && (year % 100 != 0)) || (year % 400 == 0) {
    fmt.Println("閏年です。")
  } else {
    fmt.Println("平年です。")
  }
}