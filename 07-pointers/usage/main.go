package main

import "fmt"

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(*y)
	*y = 43
	fmt.Println(*y)
}

func main() {
	x := 0

	foo(x)
	fmt.Println(x)

	fmt.Println("---")

	bar(&x)
	fmt.Println(x)
}
