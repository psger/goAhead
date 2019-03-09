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

	// 协程之间共享数据c
	go sumc(a[:len(a)/2], c) // c:17
	go sumc(a[len(a)/2:], c)  // c:-5
	x, y := <-c, <-c // 为什么结果是 -5 17

	fmt.Println(x, y, x+y)
}
