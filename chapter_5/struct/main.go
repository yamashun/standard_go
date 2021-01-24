package main

import "fmt"

type MyInt int

type (
	IntPair     [2]int
	Strings     []string
	AreaMap     map[string][2]float64
	IntsChannel chan []int
)

func main() {
	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1)
	fmt.Println(n2)

	pairt := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.7, 129.7}}

	fmt.Println(pairt)
	fmt.Println(strs)
	fmt.Println(amap)
}
