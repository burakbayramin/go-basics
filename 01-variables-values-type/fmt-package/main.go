package main

import "fmt"

var x = 42
var y = true
var z = 28

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%c\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%U\n", x)

	fmt.Printf("%t\n", y)

	s := fmt.Sprintf("%v\t%b\t%#x\t%U\n", z, z, z, z)
	fmt.Println(s)
}
