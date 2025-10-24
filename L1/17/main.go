package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7

	binarySearch(arr, target)
}

// binarySearch - функция, где внутри использует встроенный метод slices.BinarySearch
func binarySearch(arr []int, target int) int {
	index, found := slices.BinarySearch(arr, target)
	if found {
		fmt.Println(index)
		return index
	}
	fmt.Println("-1")
	return -1
}
