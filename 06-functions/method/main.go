package main

import "fmt"

type person struct {
	name string
	age  int
}

type agent struct {
	person
	gun bool
}

func (a agent) speak() {
	fmt.Println("I am", a.name)
}

func main() {
	ag := agent{
		person: person{
			name: "James Bond",
			age:  32,
		},
		gun: true,
	}
	fmt.Println(ag)
	ag.speak()

}
