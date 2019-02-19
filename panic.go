package main

import "fmt"

func main() {
	defer func() { // 在defer中通过recover来捕获panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // err就是Panic抛出的异常内容
		}
		fmt.Println("d")
	}()
	fm()
}

func fm() {
	fmt.Println("a")
	panic(55) // 抛出一个panic的异常
	fmt.Println("b")
	fmt.Println("f")
}
