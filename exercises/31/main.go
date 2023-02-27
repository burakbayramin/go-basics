package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favColors []string
}

func main() {
	p1 := person{
		firstName: "Harry",
		lastName:  "Potter",
		favColors: []string{"blue", "green", "yellow"},
	}

	p2 := person{
		firstName: "David",
		lastName:  "Cronenberg...!!",
		favColors: []string{"red", "metal", "black"},
	}

	for _, v := range p1.favColors {
		fmt.Println(v)
	}

	for _, v := range p2.favColors {
		fmt.Println(v)
	}
}
