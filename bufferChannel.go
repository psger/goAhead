package main

import "fmt"

func main() {
	//c := make(chan int, 1)
	//c <- 1
	//fmt.Println(<-c)
	//c <- 2
	//fmt.Println(<-c)

	c := make(chan int, 10)
	go fibo(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibo(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}
