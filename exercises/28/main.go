package main

import "fmt"

func main() {
	m := map[string][]string{
		"harry": {"a", "b", "c"},
		"james": {"g", "h", "j"},
		"vold":  {"n", "k", "l"},
	}

	for k, v := range m {
		fmt.Println("The record ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
