package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	//	p.name = "Agent 47"
	(*p).name = "Agent 47"
}

func main() {
	p := person{
		name: "James Bond",
	}

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
