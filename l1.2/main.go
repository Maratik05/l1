package main

import (
	"fmt"
	"sync"
)

func main() {
	ints := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for i := 0; i < len(ints); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(ints[i] * ints[i])
		}()
	}
	wg.Wait()
}
