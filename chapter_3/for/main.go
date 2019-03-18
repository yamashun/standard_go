package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 100 {
			break
		}
	}

	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	for i, r := range "ABC" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}
}
