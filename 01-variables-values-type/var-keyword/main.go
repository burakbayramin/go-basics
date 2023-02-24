package main

import "fmt"

var y = 43 //package scope

var z int

// x := 42 => ERROR

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	foo()
}

func foo() {
	fmt.Println("foo: ", y)
}
