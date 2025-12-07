package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	new_str := reverseWord(str)
	fmt.Print(new_str)
}

func reverseWord(str string) string {
	arr_word := strings.Split(str, " ")
	var rev_arr []string
	for i := len(arr_word) - 1; i >= 0; i-- {
		rev_arr = append(rev_arr, arr_word[i])
	}
	new_str := strings.Join(rev_arr, " ")
	return new_str
}
