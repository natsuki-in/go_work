package main

import (
  "fmt"
  "math/rand"
  "time"
)

type Account struct {
    mail string
    tel int
    name string
}

var letter = []rune("abcdefghijklmnopqrstuvwxyz")

func account() {  
  rand.Seed(time.Now().UnixNano())
  account := make([]Account, 3)
  fmt.Println(account)
    account[i].mail = randString(rand.Intn(10)) + "@infocom.co.jp"
    account[i].tel  = rand.Intn(9999999999)
    account[i].name = randString(rand.Intn(10))

    fmt.Printf("%v番目\n", i+1)
    fmt.Printf("メール：%v\n", account[i].mail)
    fmt.Printf("電話：%v\n", account[i].tel)
    fmt.Printf("氏名：%v\n", account[i].name) 
}

func randString (x int) string {
  str := make([]rune, x)
  for i := range str {
    str[i] = letter[rand.Intn(len(letter))]
  }
  return string(str)
}