package main

import "fmt"

var g func()

func main() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
		fmt.Println("done")
	}

	f()
	g = f
	g()
}
