package main

import "fmt"

func main() {
	x := 42

	if x == 41 {
		fmt.Println("our value was 41")
	} else if x == 42 {
		fmt.Println("our value was 42")
	} else {
		fmt.Println("our value was not 41 or 42")
	}
}
