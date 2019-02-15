package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name:"passenger"})
	fmt.Println(&person{name:"ann", age:33})

	s := person{name:"Alice", age:19}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 30
	fmt.Println("sp.age:", sp.age)
	fmt.Println("s.age:", s.age)
}
