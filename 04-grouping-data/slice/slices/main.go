package main

import "fmt"

func main() {
	x := []int{3, 4, 5, 66, 7}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[4])

	for i, v := range x {
		println(i, v)
	}

	s1 := []string{"a", "b", "c", "d"}
	s2 := []string{"f", "e", "g"}
	md := [][]string{s1, s2}
	fmt.Println(md)

}
