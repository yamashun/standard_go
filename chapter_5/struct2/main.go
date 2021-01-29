package main

import "fmt"

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed
}

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}

func main() {
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a.Name)
	fmt.Println(a.Feed.Name)
	fmt.Println(a.Feed.Amount)
	fmt.Println(a.Amount)

	a.Feed.Amount = 15
	fmt.Println(a.Feed.Amount)

	s := struct{ X, Y int }{X: 1, Y: 2}
	showStruct(s)
}
