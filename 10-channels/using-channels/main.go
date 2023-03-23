package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c) //send
	bar(c)    //receive

	fmt.Println("exit")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
