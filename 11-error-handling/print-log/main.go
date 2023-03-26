package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("non-file.txt")
	if err != nil {
		fmt.Println("err happend fmt", err)
		log.Println("err happend log", err)
	}
}
