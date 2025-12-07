package main

import "fmt"

func setBit(n int64, i uint, value int) int64 {
	if value == 1 {
		return n | (1 << i)
	}
	return n &^ (1 << i)
}

func main() {
	x := setBit(5, 0, 0)
	fmt.Print(x)
}
