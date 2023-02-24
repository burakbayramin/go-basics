package main

import "fmt"

var x int
var y float64
var z int8 // -128 to 127

func main() {
	x = 14
	y = 2.49
	//z = -129

	fmt.Printf("%v\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(z)
}
