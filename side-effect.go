package main

import "fmt"

func mutily(a, b int, typ *int) {
	*typ = a * b
}

func main(){
	a := 10
	b := 20
	typ := &a
	fmt.Println(typ)
	fmt.Println(*typ)
	mutily(a, b, typ)
	fmt.Println(a)
}
