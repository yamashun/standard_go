package main

import (
	f "./foo"
	. "fmt"
)




func main () {
	Println(f.Max)
	// fmt.Println(foo.internalFunc)
	Println(f.FooFunc(10))
	// fmt.Println(foo.internalFunc(10))
}

