package main

import "fmt"

func main() {
	m := map[string]int{
		"First":  12,
		"Second": 19,
		"Third":  31,
	}
	fmt.Println(m)
	fmt.Println(m["Second"])
	fmt.Println(m["Fifth"])

	v, ok := m["Harry"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Third"]; ok {
		fmt.Println("map contains", v)
	}
}
