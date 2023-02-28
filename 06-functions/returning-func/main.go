package main

import "fmt"

func main() {
	s := foo()
	fmt.Println(s)

	x := bar()
	i := x()
	fmt.Println(i)
	fmt.Println(x())
	fmt.Println(bar()())
}

func foo() string {
	s := "func foo"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
