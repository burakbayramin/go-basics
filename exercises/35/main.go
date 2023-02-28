package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 24, "reversed"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
	x, y := bar()
	fmt.Println(x)
	fmt.Println(y)
}
