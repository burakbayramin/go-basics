package main

import "fmt"

func main() {
	p := struct {
		name    string
		codes   map[int]string
		favGuns []string
	}{
		name: "James",
		codes: map[int]string{
			1: "Alpha",
			2: "Beta",
			3: "Tetha",
		},
		favGuns: []string{
			"ak47",
			"m4a1",
			"desert eagle",
		},
	}

	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.codes)
	fmt.Println(p.favGuns)

}
