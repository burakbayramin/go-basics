package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)  //value
	fmt.Println(&a) //adress of memory
	//fmt.Println(*&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)

	c := &a
	fmt.Println(c)
	fmt.Println(*c)

	i := 12
	j := &i
	*j = 13
	fmt.Println(i)
}
