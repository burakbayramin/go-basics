package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota
	e = iota + 1
	f = iota + 5
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
