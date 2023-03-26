package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("9 + 7 =", mySum(9, 7))
	fmt.Println("5 + 7 =", mySumm(5, 7))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func mySumm(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum + 1
}
