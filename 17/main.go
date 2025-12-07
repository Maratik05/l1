package main

import "fmt"

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("До сортировки:", arr)
	sorted_arr := quickSort(arr)
	fmt.Println("После сортировки:", sorted_arr)
	el := binSearch(sorted_arr, 12)
	fmt.Print(el)
}

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

func binSearch(arr []int, x int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
