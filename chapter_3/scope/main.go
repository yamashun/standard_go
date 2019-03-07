package main

import (
	"./foo"
	"fmt"
)




func main () {
	fmt.Println(foo.Max)
	// fmt.Println(foo.internalFunc)
	fmt.Println(foo.FooFunc(10))
	// fmt.Println(foo.internalFunc(10))
}

