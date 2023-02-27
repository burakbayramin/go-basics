package main

import "fmt"

type person struct {
	fname string
	age   int
}

type agent struct {
	person
	gun bool
}

func main() {
	ag := agent{
		person: person{
			fname: "James",
			age:   32,
		},
		gun: true,
	}

	p := person{
		fname: "Harry",
		age:   16,
	}

	fmt.Println(ag)
	fmt.Println(ag.fname, ag.age, ag.gun)
	fmt.Println(p)
	fmt.Println(p.fname, p.age)
}
