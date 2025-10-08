package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	var unicalStr []string
	exMap := make(map[string]bool)
	for _, str := range strs {
		if exMap[str] {
			continue
		}
		unicalStr = append(unicalStr, str)
		exMap[str] = true
	}
	fmt.Print(unicalStr)
}
