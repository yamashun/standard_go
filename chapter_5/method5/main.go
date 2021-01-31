package main

import "fmt"

type Point struct{ X, Y int }

func (p Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

func (p *Point) Set2(x, y int) {
	p.X = x
	p.Y = y
}

func main() {
	p1 := Point{}
	p1.Set(1, 2)
	fmt.Println(p1.X)
	fmt.Println(p1.Y)

	p2 := Point{}
	p2.Set(1, 2)
	fmt.Println(p2.X)
	fmt.Println(p2.Y)

	p1.Set2(1, 2)
	fmt.Println(p1.X)
	fmt.Println(p1.Y)

	p2.Set2(1, 2)
	fmt.Println(p2.X)
	fmt.Println(p2.Y)
}
