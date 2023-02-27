package main

import "fmt"

// func (r reciever) identifier(parameter(s)) (return(s)) { code }

func main() {
	foo()
	bar("James")
	s := woo("Canada")
	fmt.Println(s)
	x, y := tang("Harry", "Potter")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from, ", s)
}

func tang(first, last string) (string, bool) {
	a := fmt.Sprint("is ", first, " ", last, " agent ?")
	b := false
	return a, b
}
