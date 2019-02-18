package main

import "fmt"

type person struct {
	name string
	age int
}

type student struct {
	person
	age int
	class string
}

type teacher struct {
	person
	money float32
}

func main() {
	a := student{person:person{"passenger", 20}, class:"1501"}
	fmt.Println(a.person.age)
	a.person.age = 21
	fmt.Println(a.person.age)
	a.age = 22
	fmt.Println(a.age)
	fmt.Println(person{name:"passenger"})
	fmt.Println(&person{name:"ann", age:33})

	s := person{name:"Alice", age:19}
	fmt.Println(s.name)

	var t teacher
	t.money = 3000
	fmt.Println(t)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 30
	fmt.Println("sp.age:", sp.age)
	fmt.Println("s.age:", s.age)
}
