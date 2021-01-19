package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	gs := 100
	var wg1 sync.WaitGroup
	var incrementer int64

	wg1.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			wg1.Done()
		}()

	}
	wg1.Wait()
	fmt.Println(incrementer)
}
