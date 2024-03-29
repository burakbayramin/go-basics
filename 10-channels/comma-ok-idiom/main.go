package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)
	receive(even, odd, quit)

	fmt.Println("program finished with success")

	/////////

	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel -> ", v)
		case v := <-o:
			fmt.Println("from the odd channel -> ", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from ok -> ", ok, "---", i)
				return
			} else {
				fmt.Println("from ok (else block) -> ", ok, "---", i)
			}
		}
	}
}
