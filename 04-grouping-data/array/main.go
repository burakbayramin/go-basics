package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	a[4] = 15
	fmt.Println(a)
	fmt.Println(len(a))

	b := [4]int{23, 12, 4, 66}
	fmt.Println(b)

	var td [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			td[i][j] = i + j
		}
	}
	fmt.Println(td)
}
