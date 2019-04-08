package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}

	s, result := m[1]
	s2, result2 := m[9]

	fmt.Printf("%s:%t\n", s, result)
	fmt.Printf("%s:%t\n", s2, result2)
}
