package main

import "fmt"

func inc(p *int) {
	*p++
}

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

func main() {
	a := [3]string{"Apple", "Banana", "Cherry"}
	p := &a
	fmt.Println(a[1])
	fmt.Println(p[1])
	p[2] = "Grape"
	fmt.Println(a[2])
	fmt.Println(p[2])
}
