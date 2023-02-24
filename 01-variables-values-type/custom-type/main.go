package main

import "fmt"

var a int

type tea int

var b tea

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	//CONVERSION
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
