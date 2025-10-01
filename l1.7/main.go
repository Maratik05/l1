package main

import (
	"fmt"
	"strconv"
	"sync"
)

func writeOnMap(x int, m map[int]string, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	mtx.Lock()
	m[x] = strconv.Itoa(x)
	mtx.Unlock()

}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	m := make(map[int]string, 10)
	for i := 1; i < 110000; i++ {
		wg.Add(1)
		go writeOnMap(i, m, &wg, &mtx)
	}
	wg.Wait()
	fmt.Print(m)
}
