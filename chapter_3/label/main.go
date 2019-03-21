package main

import (
	"fmt"
)

func main() {
LOOP:
	for {
		for {
			for {
				fmt.Println("start")
				break LOOP
			}
			fmt.Println("ここは通らない")
		}
		fmt.Println("ここは通らない")
	}
	fmt.Println("完了")
}
