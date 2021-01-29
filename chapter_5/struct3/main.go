package main

import "fmt"

type Point struct {
	X, Y int
}

func swap(p Point) {
	x, y := p.X, p.Y
	p.X = y
	p.Y = x
}

func main() {
	p := Point{X: 1, Y: 2}
	swap(p)

	fmt.Println(p.X)
	fmt.Println(p.Y)
}
