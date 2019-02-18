package main

import (
	"fmt"
	"reflect"
)

func main()  {
	x := 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
