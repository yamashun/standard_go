package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5)

	fmt.Println(s)

	// 要素数を超えて要素へアクセスしうとするとランタイムパニック
	// fmt.Println(s[5])

	fmt.Println(len(s))
	fmt.Println(cap(s))

	s2 := make([]int, 5, 10)

	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
