package main

import (
  "fmt"
)

func main() {
	str := "else if"

	if str == "if" {
		fmt.Println("if")
	} else if str == "else if" {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	// 簡易文つきif
	if x, y := 1, 2; x < y {
		fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	}
}
