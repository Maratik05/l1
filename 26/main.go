package main

import "fmt"

func main() {
	str := "ab"
	a := isUnique(str)
	fmt.Print(a)
}

func isUnique(str string) bool {
	symb_exist := make(map[rune]bool)
	for _, el := range str {
		value, ok := symb_exist[rune(el)]
		if !ok {
			symb_exist[rune(el)] = true
		}
		if value == true {
			return false
		}
	}
	return true
}
