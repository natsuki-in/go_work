package main

import "fmt"

func multipulOf7() {
  // 0から100までの整数
  for i := 0; i <= 100 ; i++{
    // 7の倍数のみ表示
    if i != 0 && i % 7 == 0 {
      fmt.Println(i, "は７の倍数です。")
    }
  }
}