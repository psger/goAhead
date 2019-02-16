package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := make([]string, 3) //
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("slices s:", l)

	l = s[:5]
	fmt.Println("slice2:", l)

	l = s[2:]
	fmt.Println("slice3", l)
	var m = []string{"你好", "你", "好", "啊"}
	m = append(m, "i")
	fmt.Println("m:", m[:2])
	fmt.Println(reflect.ValueOf(m).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())

	b := [8]int{1, 2, 3, 4, 5} // 数组和slice的区别 数组为定长
	fmt.Println(reflect.ValueOf(b).Kind())
	slic2 := make([]int, 3)
	slic2[0] = 12
	fmt.Println(slic2)

	newSlice := make([]int, 10, 20)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))

	var arr = [5]int{1, 4, 5, 6, 2}
	fmt.Println(reflect.ValueOf(arr).Kind()) // array
	fmt.Println(reflect.ValueOf(arr[:]).Kind()) //slice
	//s1 := []byte{'p', 'o', 'e', 'm'} 且 s2 := d[2:]，s2 的值是多少？如果我们执行 s2[1] == 't'，s1 和 s2 现在的值又分配是多少？

	s1 := []byte{'p', 'o', 'e', 'm'}
	fmt.Println(len(s1[2:2]))
	fmt.Println(s1)
	s2 := s1[2:]
	fmt.Println(s2)
	s2[1] = 't'
	fmt.Println(s2)
	fmt.Println(s1)
}
