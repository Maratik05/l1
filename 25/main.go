package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	sleep(5 * time.Second)
	fmt.Println(time.Now())
}

func sleep(duration time.Duration) {
	done := make(chan bool)

	time.AfterFunc(duration, func() {
		done <- true
	})
	<-done
}
