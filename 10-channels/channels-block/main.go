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
}
