package main

import (
	"fmt"
)

func pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func pow2(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func main() {
	// s := make([]int, 5)

	// fmt.Println(s)

	// // 要素数を超えて要素へアクセスしうとするとランタイムパニック
	// // fmt.Println(s[5])

	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// s2 := make([]int, 5, 10)

	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))

	// // 管理スライス式
	// a := [5]int{1, 2, 3, 4, 5}
	// s3 := a[0:2]
	// fmt.Println(s3)

	// // append
	// s = append(s, 1)
	// fmt.Println(s)
	// // 複数の要素追加
	// s = append(s, 2, 3, 4)
	// fmt.Println(s)
	// // sliceの追加
	// s = append(s, []int{1, 2}...)
	// fmt.Println(s)

	// 配列は値渡し
	// a:= [3]int{1, 2, 3}
	// pow(a)
	// fmt.Println(a)

	// // sliceは参照渡し
	// b:= []int{1, 2, 3}
	// pow2(b)
	// fmt.Println(b)

	// a := [5]int{1, 2, 3, 4, 5}
	// s := a[0:2]
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// a[1] = 0
	// fmt.Println(s)

	a := [3]int{1, 2, 3}
	s := a[:]
	fmt.Println(s)

	// appendで自動拡張されない場合は、もとの配列を参照したまま
	x := a[0:2]
	x = append(x, 4)
	a[1] = 0
	fmt.Println(s)

	// appendで自動拡張される場合、元の配列とは異なる新たに確保されたメモリ領域を参照
	y := append(s, 4)
	a[0] = 9
	fmt.Println(y)
}
