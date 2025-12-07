package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func writeToChannel(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	var num int
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Scan(&num)
			select {
			case ch <- num:
			case <-ctx.Done():
				return
			}
		}
	}
}

func readFromChannel(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case num, ok := <-ch:
			if !ok {
				fmt.Print("channel closed!!!")
				return
			}
			fmt.Printf("Readed from channel number: %d\n", num)

		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go writeToChannel(ctx, ch, &wg)
	go readFromChannel(ctx, ch, &wg)

	wg.Wait()
}
