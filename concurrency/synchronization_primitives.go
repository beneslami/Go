package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // 1- first step is to initialize in package scope

func main(){
	fmt.Println("OS\t", runtime.GOOS);
	fmt.Println("ARCH\t", runtime.GOARCH);
	fmt.Println("CPUs\t", runtime.NumCPU());
	fmt.Println("GoRoutines\t", runtime.NumGoroutine());

	wg.Add(1); // 2- second step is to add an item to the queue

	go test1();  // concurrency
	test2();

	wg.Wait(); // 3- third step is to wait for the concurrent process to finish
}

func test1(){
	for i := 0; i < 10; i++ {
		fmt.Println(i);
	}
	wg.Done(); // 4- forth step is to dequeue
}

func test2(){
	for i := 0; i < 10; i++ {
		fmt.Println(i);
	}
}
