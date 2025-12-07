package main

import (
	"fmt"
)

func reverseString(str string) string {
	var runes []rune = []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	new_str := string(runes)
	return new_str
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	rev_str := reverseString(str)
	fmt.Print(rev_str)
}
