package main

import "fmt"

func foo(xi []int) int {
	return xi[0] + xi[len(xi)-1]
}

func bar(f func(xi []int) int, ii []int) int {
	return f(ii)
}

func main() {
	fmt.Println(bar(foo, []int{1, 2, 3, 4, 5}))
}
