package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func (r *rect) add() {
	r.height = 100
	fmt.Println(r)
}

//type intr interface {
//	area() int
//	perim() int
//	add()
//}


func main()  {
	r := rect{width:10, height:5}
	//var inte rect
	//inte = r
	//fmt.Println(inte.perim())
	//fmt.Println(inte.area())
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("prim:", rp.perim())
}