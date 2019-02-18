package main

import (
	"fmt"
	"reflect"
)

//
//type Point struct {
//	x, y float64
//}
//
//func (p *Point) Abs() float64 {
//	return math.Sqrt(p.x*p.x + p.y*p.y)
//}
//
//type NamedPoint struct {
//	Point
//	name string
//}
//
//func main() {
//	//var n *NamedPoint
//	n := NamedPoint{Point{3, 4}, "Pythagoras"}
//	fmt.Println(n.Abs()) // 打印5
//}

//type Shaper interface {
//	Area() float32
//	tim() string
//}
//
//type Square struct {
//	side float32
//}
//
//func (sq *Square) Area() float32 {
//	return sq.side * sq.side
//}
//
//func main() {
//	//sq1 := new(Square)
//	//sq1.side = 5
//
//	var sq1 Square
//	sq1.side = 5
//	fmt.Println(sq1.Area())
//
//	// var areaIntf Shaper
//	// areaIntf = sq1
//	// shorter,without separate declaration:
//	// areaIntf := Shaper(sq1)
//	// or even:
//	areaIntf := sq1
//	fmt.Println(areaIntf)
//	fmt.Printf("The square has area: %f\n", areaIntf.Area())
//	fmt.Println(areaIntf)
//}

type user struct {
	name  string
	email string
}
type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n",
		u.name,
		u.email)
}
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{
		name:  "stormzhu",
		email: "abc@qq.com",
	}
	//sendNotification(u)

	var send notifier
	send := u
	send.notify()
	fmt.Println(reflect.TypeOf(send).Kind())
}