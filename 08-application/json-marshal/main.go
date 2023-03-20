package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Luke",
		Last:  "Skywalker",
		Age:   32,
	}

	p2 := person{
		First: "Anakin",
		Last:  "Skywalker",
		Age:   60,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, er := json.Marshal(people)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(bs))
}
