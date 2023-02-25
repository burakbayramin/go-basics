package main

import "fmt"

func main() {
	i := 2000
	for {
		if i == 2024 {
			break
		}
		fmt.Println(i)
		i++
	}
}
