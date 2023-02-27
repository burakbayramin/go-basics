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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
