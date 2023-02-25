package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m)
	m["e"] = 4
	fmt.Println(m)

	delete(m, "e")
	fmt.Println(m)

	if v, ok := m["b"]; ok {
		fmt.Println("value is ", v)
		delete(m, "b")
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
