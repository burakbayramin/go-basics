package main

import "fmt"

func main() {

	if true {
		fmt.Println("first statement")
	}
	if false {
		fmt.Println("second statement")
	}
	if !true {
		fmt.Println("third statement")
	}
	if !false {
		fmt.Println("fourth statement")
	}

}
