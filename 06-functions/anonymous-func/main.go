package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("anonymous func")
	}()

	func(x int) {
		fmt.Println("meaning of life:", x)
	}(42)
}

func foo() {
	fmt.Println("normal func")
}
