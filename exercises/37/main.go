package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer func")
	}()
	fmt.Println("hello")
}
