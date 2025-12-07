package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	v  int
	mu sync.Mutex
}

func (c *Counter) increment() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}
func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}
	wg.Wait()
	fmt.Print(counter.v)
}
