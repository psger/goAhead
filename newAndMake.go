package main

import (
	"fmt"
	"reflect"
)

func main()  {
	v := new(int)
	fmt.Println(v)
	fmt.Println(reflect.ValueOf(v).Kind())

	m := make([]string, 2)
	fmt.Println(m)
	fmt.Println(reflect.ValueOf(m).Kind())

	a := 12
	fmt.Println(&a)
	var pt *int
	pt = &a
	fmt.Println(*pt)

	var slice1 []string
	fmt.Println(slice1)
	fmt.Println(reflect.ValueOf(slice1).Kind())


}
