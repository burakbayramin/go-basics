package main

import "fmt"

func main() {
	favSport := "boxing"
	switch favSport {
	case "tennis":
		fmt.Println("not fav")
	case "baseball":
		fmt.Println("not fav")
	case "boxing":
		fmt.Println("fav sport")
	}
}
