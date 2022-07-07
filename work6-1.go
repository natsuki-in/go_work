package main

import "fmt"

func check_even(a int) {
  n := a
  if (n % 2 == 0) {
    fmt.Println("偶数です。")
  } else {
    fmt.Println("奇数です。")
  }
}