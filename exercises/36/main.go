package main

import "fmt"

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(yi []int) int {
	total := 0
	for _, v := range yi {
		total += v
	}
	return total
}

func main() {
	i := []int{1, 2, 3, 4, 5}
	ii := foo(i...)
	fmt.Println(ii)

	j := []int{1, 2, 3, 4, 5}
	jj := bar(j)
	fmt.Println(jj)
}
