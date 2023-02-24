package main

import "fmt"

func main() {
	s := "Abrakadabra"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%q", s[i])
	}

	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

}
