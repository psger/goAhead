package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 3)
	fmt.Println(res)

	res = plusPlus(2, 4, 6)
	fmt.Println(res)
}