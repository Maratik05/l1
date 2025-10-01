package main

import (
	"fmt"
	"sync"
)

func writeInChannel(ch chan<- int, el int) {
	ch <- el

}

func writeInChannelDouble(ch <-chan int) chan int {
	dch := make(chan int)
	go func() {
		defer close(dch)
		for el := range ch {
			dch <- el * 2
		}
	}()

	return dch
}

func main() {
	elch := make(chan int)
	dch := writeInChannelDouble(elch)
	var wg sync.WaitGroup
	var arr []int = []int{2, 4, 7, 5, 4, 3, 2, 0, 9, 8, 7}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(elch)
		for _, el := range arr {

			writeInChannel(elch, el)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for el := range dch {
			fmt.Println(el)
		}
	}()
	wg.Wait()
}
