package main

import "fmt"

func main() {
	arr := [5]int{0, 1, 2, 3, 4}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
}
