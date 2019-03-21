package main

import (
	"fmt"
)

func main() {
	s := "A"
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
		fallthrough
	case "C":
		s += "D"
		fallthrough
	default:
		s += "E"
	}
	fmt.Println(s)

	switch n :=2; n {
	case 1, 3, 5, 7, 9:
		fmt.Printf("%d is odd\n", n)
	case 2, 4, 6, 8, 10:
		fmt.Printf("%d is even\n", n)
	}

	var x interface{} = 3
	// i := x.(int)
	// f := x.(float64)
	// 2つの変数を代入するような形式で書くことで2番目の変数にアサーションの結果が入る

	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("integer or unsigned integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("don't know")
	}
}

