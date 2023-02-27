package main

import "fmt"

func main() {

	p := struct {
		fname string
		age   int
	}{
		fname: "David",
		age:   41,
	}

	fmt.Println(p.fname, p.age)
}
