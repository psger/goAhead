package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 将closure赋值给变量
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) {fmt.Println(i)}
		// g(i)
		// g的值是一个内存地址
		fmt.Println(g)
	}
}

// closure能够保存并积累其中的变量 不论外部函数退出与否
func adder1() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}


func main() {
	//nextInt := intSeq()
	//
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//
	//newInts := intSeq()
	//fmt.Println(newInts())

	//f()
	f := adder1()
	fmt.Print(f(1))
	fmt.Print(f(12))
	fmt.Println(f(24))
}