package main

import "fmt"

func main() {
	f1 := func() {
		fmt.Println("func expression")
	}

	f2 := func(x int) {
		fmt.Println("meaning of life:", x)
	}

	f1()
	f2(42)
}
