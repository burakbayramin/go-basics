package main

import "fmt"

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4, 5, 6}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	fmt.Printf("%T\n", nums)

	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}
