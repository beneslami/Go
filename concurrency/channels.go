package main

import "fmt"

func main() {
	/*  1-
	c := make(chan int)       // We can also eliminate the go func() section by adding c := make(chan int 1)
	defer close(c)
	go func(){
		c <- 42
	}()
	fmt.Println(<-c)
	*/

	/*  2-
	cs := make(chan <- int)  // Send-Only channel. Similar to this, make(<- chan int) is receive-only channel
	go func(){
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Println("------\n")
	fmt.Printf("cs\t%T\n", cs)
	*/

	/* 3-
	c := gen()
	receive(c)
	fmt.Println("about to exit")
	*/
	/* 4-
	q := make(chan int)
	c := gen(q)
	receive(c, q)
	fmt.Println("about to exit")
	*/

	/* 5-
	c := make(chan int)

	go func(){
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
	*/

	c := make(chan int, 100)
	value := 0
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- value
				value++
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}

}

/* 3- functions
func gen() <-chan int {
	c := make (chan int)
	go func(){
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int){
	for v := range c {
		fmt.Println(v)
	}
}
*/

/* 4-
func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}

func receive(c, q <-chan int){
	for{
		select{
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
*/
