package main

import "fmt"

type person struct {
	name string
	age  int
}

type soldier struct {
	person
	strength int
	gun      bool
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("just a person")
	case soldier:
		fmt.Println("i can be your superman")

	}
	fmt.Println("passed into bar ", h)
}

func (p person) speak() {
	fmt.Println("I am ", p.name, " -person")
}

func (s soldier) speak() {
	fmt.Println("I am ", s.name, " -soldier")
}

func main() {
	prsn := person{
		name: "Jeff",
		age:  31,
	}

	sldr := soldier{
		person: person{
			name: "Matt",
			age:  32,
		},
		strength: 86,
		gun:      true,
	}

	prsn.speak()
	sldr.speak()
	fmt.Printf("%T\n", prsn)
	fmt.Printf("%T\n", sldr)
	bar(prsn)
	bar(sldr)

}
