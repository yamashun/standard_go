package main

import (
  "fmt"
)

func plus(x, y int) int {
  return x + y
}

var plusAlias = plus

func returnFunc() func() {
  return func() {
    fmt.Println("I'm a function")
  }
}

func main() {
  fmt.Printf("%v\n", func(x, y int) int { return x + y })
  fmt.Printf("%v\n", func(x, y int) int { return x + y }(2, 3))

  fmt.Println(plusAlias(10, 5))

  // 変数に代入して呼び出す
  f := returnFunc()
  f()

  // 変数を経由せずに呼び出す
  returnFunc()()
}
