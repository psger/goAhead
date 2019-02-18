package main

import "fmt"

func sumc(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // 发送total到channel c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sumc(a[:len(a)/2], c)
	go sumc(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
