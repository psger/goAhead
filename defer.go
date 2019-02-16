package main

import "fmt"

func main() {
	function1()
	a()
	fmt.Println(a())
	doDBOptions()
}

func function1() {
	fmt.Println("111")
	defer function2()
	fmt.Println("33333")
}

func function2() {
	fmt.Println("22222")
}

func a() int {
	i := 0
	defer fmt.Println(i)
	i++
	return i
}

func connToDB() {
	fmt.Println("connected to Db")
}

func disconnFromDb() {
	 fmt.Println("disconnected from DB")
}

func doDBOptions() {
	connToDB()
	fmt.Println("延迟关闭连接")
	defer disconnFromDb()
	fmt.Println("一些数据库操作")
}