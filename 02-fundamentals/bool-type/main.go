package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	fmt.Println(!x)

	a := 41
	b := 42
	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a != b)

}