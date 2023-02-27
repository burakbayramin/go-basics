package main

import "fmt"

func main() {
	m := map[string][]string{
		"harry": {"a", "b", "c"},
		"james": {"g", "h", "j"},
		"vold":  {"n", "k", "l"},
	}

	delete(m, "james")

	for i, v := range m {
		fmt.Println("record of ", i)
		for j, v2 := range v {
			fmt.Println("\t", j, v2)
		}
	}
}
