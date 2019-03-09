package main

import "fmt"

func partation(array []int, left, right int) int {
	key := array[left]
	index := left + 1
	for i := left + 1; i <= right; i++{
		if array[i] < key {
			swap(array, i, index)
			index++
		}
	}
	swap(array, left, index)
	return index - 1
}

func swap(array []int, i, j int) {
	var tmp int
	tmp = array[i]
	array[i] = array[j]
	array[j] = tmp
}

func quickSort(array []int, left, right int) {
	if left < right {
		mid := partation(array, left, right)
		partation(array, left, mid - 1)
		partation(array, mid + 1, right)
	}
}

func main() {
	array := []int{4,5,7,1,2,9,6,3}
	quickSort(array, 0, 7)
	fmt.Println(array)
}