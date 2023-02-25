package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 28
	x[9] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//x[10] = 43 ERROR
	x = append(x, 43)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
