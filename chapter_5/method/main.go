package main

import (
	"fmt"
	"math"
)

type IntPoint struct{ X, Y int }
type FloatPoint struct{ X, Y float64 }

func (p *IntPoint) Distance(dp *IntPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (p *FloatPoint) Distance(dp *FloatPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(x*x + y*y)
}

func main() {
	p := &IntPoint{X: 0, Y: 0}
	result := p.Distance(&IntPoint{X: 1, Y: 1})
	fmt.Println(result)
}
