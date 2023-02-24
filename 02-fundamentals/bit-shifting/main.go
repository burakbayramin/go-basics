package main

import "fmt"

const (
	_  = iota
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
)

func main() {
	x := 4
	y := x << 1
	fmt.Printf("%d\t%b\n", x, x)
	fmt.Printf("%d\t%b\n", y, y)

	fmt.Println()

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
