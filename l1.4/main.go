package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d завершает работу\n", id)
			return

		default:
			fmt.Printf("Worker %d выполняет работу!!!!!", id)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	cancel()
	wg.Wait()

}
