package main

import "fmt"

func main() {
	sl1 := []string{"a", "b", "c"}
	sl2 := []string{"d", "e", "f"}
	md := [][]string{sl1, sl2}

	for i, v := range md {
		fmt.Println("Record ", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v \t value: %v\n", j, val)
		}
	}
}
