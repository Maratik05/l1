package main

import "fmt"

func deleteAtIndex(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}

	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("До:", numbers)

	numbers = deleteAtIndex(numbers, 2)
	fmt.Println("После:", numbers)

	fmt.Printf("Длина: %d, ёмкость: %d\n", len(numbers), cap(numbers))
}
