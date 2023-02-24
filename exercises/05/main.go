package main

import "fmt"

type cat int

var x cat
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v\n%T\n", y, y)
}
