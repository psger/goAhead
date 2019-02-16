package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	fmt.Println("pointer:", &i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	a := 20
	var ap *int //声明一个指针

	ap = &a // 指针变量的存储地址
	fmt.Println("a变量的地址是:", &a)
	fmt.Println("ap变量存储的指针地址：", ap)
	fmt.Println("使用指针访问值：", *ap)
}
