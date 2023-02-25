package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		fmt.Println("outer loop: ", i)
		for j := 0; j < 3; j++ {
			fmt.Println("\tinner loop: ", j)
		}
	}

	for i := 0; i < 6; i++ {
		fmt.Println()
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
	}

	fmt.Println()
	fmt.Println()


	for i := 1; i < 10; i += 2 {
		for k := 0; k < (4 - i/2); k++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
