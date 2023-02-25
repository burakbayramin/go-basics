package main

import "fmt"

func main() {
	x := []int{45, 66, 12, 33, 98, 56, 3}
	fmt.Println(x)
	fmt.Println(x[:1])
	fmt.Println(x[1:])
	fmt.Println(x[2:5])
	fmt.Println(x[1:5])
	
}