package main

import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(animals.ElephantFood())
	fmt.Println(animals.MonkeyFood())
	fmt.Println(animals.RabbitFood())
}
