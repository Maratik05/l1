package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	curr_el := arr[0]
	var left, right []int
	for _, value := range arr[1:] {
		if value < curr_el {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	return append(append(quickSort(left), curr_el), quickSort(right)...)
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("До сортировки:", arr)
	sorted_arr := quickSort(arr)
	fmt.Println("После сортировки:", sorted_arr)
}
