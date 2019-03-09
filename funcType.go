package main

import "fmt"

type interFunc func(int) (bool)

func isOdd(integer int) (bool) {
	if integer % 2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f interFunc) []int {
	var result []int
	for _, s := range slice {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}

func main() {
	s := []int{1,2,3,4,5}
	res := filter(s, isOdd)
	fmt.Println(res)
}