package main

import (
	"fmt"
	"pack1"
)

func main() {
	fmt.Println("hello world")
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Println(test1)
}