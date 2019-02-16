package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func fibonacci(n int) (res int) {
	if n < 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}

	return
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(fibonacci(20))
}