package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	///////

	cc := make(chan int, 1)

	cc <- 42

	fmt.Println(<-cc)

	////////

	ch := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	fmt.Printf("%T\n", ch)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
