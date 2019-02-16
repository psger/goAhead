package main

import "fmt"

// add2返回一个匿名函数
func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

// adder返回一个匿名函数
func adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func main()  {
	p2 := add2()
	fmt.Println(p2(3))
	twoAdder := adder(4)
	fmt.Println(twoAdder(6))
}
