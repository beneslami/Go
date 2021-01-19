package main

import (
	"fmt"
	"sync"
)

func main() {
	incrementer := 0
	gs := 100
	var wg1 sync.WaitGroup
	var mutex = sync.Mutex{}

	wg1.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := incrementer
			v++
			incrementer = v
			mutex.Unlock()
			wg1.Done()
		}()

	}
	wg1.Wait()
	fmt.Println(incrementer)
}
