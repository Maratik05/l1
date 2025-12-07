package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func gorutine1(s chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-s:
			fmt.Println("Finished gorutine")
			return
		default:
			fmt.Println("Working gorutine")

		}
	}
}

func gorutine2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Finished gorutine2")
			return
		default:
			fmt.Println("Working gorutine2")
		}
	}
}

func gorutine3(s chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-s:
			fmt.Println("Finished gorutine3")
			runtime.Goexit()
		default:
			fmt.Println("Working gorutine3")
		}
	}

}

func main() {
	ch := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(3)
	go gorutine1(ch, &wg)
	go gorutine2(ctx, &wg)
	go gorutine3(ch, &wg)

	time.Sleep(10 * time.Second)
	cancel()
	ch <- true
	ch <- true
	wg.Wait()
}
