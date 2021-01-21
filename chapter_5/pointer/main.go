package main

import "fmt"

func main() {
	var i int
	p := &i
	i = 5

	fmt.Println(*p)

	*p = 10
	fmt.Println(i)
}
