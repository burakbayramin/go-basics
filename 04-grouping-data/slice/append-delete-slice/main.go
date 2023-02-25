package main

import "fmt"

func main() {
	//APPEND
	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	y := []int{44, 5, 77}
	x = append(x, y...)
	fmt.Println(x)

	//DELETE
	x = append(x[:2], x[5:]...)
	fmt.Println(x)
}
