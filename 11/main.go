package main

import (
	"fmt"
)

func intersection(arr1 []int, arr2 []int) []int {
	var res []int
	existMap := make(map[int]bool)
	for _, el := range arr1 {
		existMap[el] = true
	}

	for _, el := range arr2 {
		if existMap[el] {
			res = append(res, el)
		}
	}
	return res
}
func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 3, 4}
	fmt.Print(intersection(arr1, arr2))
}
