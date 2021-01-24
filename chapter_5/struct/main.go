package main

import "fmt"

type MyInt int

func main() {
	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1)
	fmt.Println(n2)
}
