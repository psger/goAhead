package main

import "fmt"

func main() {
	x := Min(1, 9, 19, 3, 0)
	fmt.Println(x)
	arr := []int{4,5,6,7}
	x = Min(arr...)
	fmt.Println(x)

}

func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
