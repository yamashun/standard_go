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
	i := x.(int)
	// f := x.(float64)
	_,isFloat64 := x.(float64)

	fmt.Println(i)
	// fmt.Println(f)
	fmt.Println(isFloat64)
}

