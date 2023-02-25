package main

import "fmt"

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println()

	j := 2

	switch j {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	}

	fmt.Println()

	n := "Bond"

	switch n {
	case "Bond":
		fmt.Println("My name is bond, james bond")
	case "Harry":
		fmt.Println("My name is haarrry potttaa")
	default:
		fmt.Println("default")
	}
}
