package main

import (
	"fmt"
	"runtime"
	"time"
)

func fibonacciNew(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-time.After(5 * time.Second):
			println("timeOut")
			break
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// 定义多个channel
	c := make(chan int)
	quit := make(chan int)
	fmt.Println(runtime.NumGoroutine())
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fmt.Println(runtime.NumCPU())
	//fibonacciNew(c, quit)
}
