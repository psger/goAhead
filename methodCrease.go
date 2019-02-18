package main

import "fmt"

type TZ int

func (tz *TZ) append(num int) {
	*tz += TZ(num)
}

func main() {
	var a TZ
	a.append(100)
	fmt.Println(a)
}