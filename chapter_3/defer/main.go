package main

import (
	"fmt"
)

func main() {
	runDefer()
}

func runDefer() {
	defer fmt.Println("defer")
	fmt.Println("done")
}
