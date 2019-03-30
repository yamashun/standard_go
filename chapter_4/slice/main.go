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

	// 管理スライス式
	a := [5]int{1, 2, 3, 4, 5}
	s3 := a[0:2]
	fmt.Println(s3)

	// append
	s = append(s, 1)
	fmt.Println(s)
}
