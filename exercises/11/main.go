package main

import "fmt"

const (
	a = 2023 + iota
	b = 2023 + iota
	c = 2023 + iota
	d = 2023 + iota
)

func main() {
	fmt.Println(a,b,c,d)
}