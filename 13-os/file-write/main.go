package main

import (
	"fmt"
	"os"
)

func main() {
	writeOsWrite()
	writeOsCrate()
}

func writeOsWrite() {
	err := os.WriteFile("testfile1.txt", []byte("test"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func writeOsCrate() {
	f, err := os.Create("testfile2.txt")
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte("1\n"))
	f.Write([]byte("2\n"))
	f.Write([]byte("3\n"))
	f.Write([]byte("James\n"))
	f.Write([]byte("Bond\n"))
	f.Close()
}
