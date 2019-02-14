package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["k1"] = "2"
	m["k2"] = "3"

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	k2, prs := m["k2"] //检查是否存在该key的值 存在prs:true, k1:value 不存在prs:false k1:(int->0, string->"")
	fmt.Println("prs:", prs)
	fmt.Println("k1:", k2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	var mn map[int]map[int]string
	mn = make(map[int]map[int]string)
	a, ok := mn[2][1]
	if !ok {
		mn[2] = make(map[int]string)
	}
	mn[2][1] = "MeiM"
	a = mn[2][1]

	fmt.Println(a, ok)
}
