package main

import "fmt"

func worker(ch <-chan int, n int) {

	for b := range ch {
		fmt.Println("worker number", n, "read", b)
	}
}

func main() {
	var n int
	ch := make(chan int)
	x := 1
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		go worker(ch, i)
	}

	for {
		ch <- x
	}
}
