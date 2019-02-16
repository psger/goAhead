package main

import "fmt"

func main() {
	callBack(3, Add)
}

func Add(a, b int) {
	fmt.Println(a, b, a+b)
}

func callBack(y int, f func(int, int)) {
	f(y, 2)
}
