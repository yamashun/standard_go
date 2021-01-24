package main

import "fmt"

type MyInt int

type (
	IntPair     [2]int
	Strings     []string
	AreaMap     map[string][2]float64
	IntsChannel chan []int
)

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

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

	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)

	fmt.Println(n)
}
