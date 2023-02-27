package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {

	p1 := person{
		fname: "James",
		lname: "Bond",
	}

	p2 := person{
		fname: "Harry",
		lname: "Potter",
	}

	fmt.Println(p1)
	fmt.Println(p1.fname)
	fmt.Println(p1.lname)
	fmt.Println(p2.fname, p2.lname)
}
