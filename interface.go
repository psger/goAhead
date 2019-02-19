package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}
//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
// Human对象实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Employee重载Human的Sayhi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

func (e Employee) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}


type Men interface {
	SayHi()
	Sing(lyrics string)
	//Guzzle(beerStein string)
}



func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	//paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	//sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	var interf Men // 定义Men类型的变量
	interf = mike
	fmt.Println(reflect.TypeOf(interf))
	//i能存储Student
	i := &mike
	fmt.Println(reflect.TypeOf(i).Kind())
	fmt.Println("This is Mike, a Student:")
	i.SayHi()

	i.Sing("November rain")

	//i也能存储Employee
	im := tom
	fmt.Println("This is tom, an Employee:")
	im.SayHi()
	im.Sing("Born to be wild")
}



