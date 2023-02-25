package main

import "fmt"

func main() {

	for i := 0; i < 25; i++ {
		if i == 16 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
