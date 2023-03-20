package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Runtimes\t", runtime.NumGoroutine())

	wg.Add(1)

	go foo()
	bar()

	fmt.Println("Runtimes\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo\t", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar\t", i)
	}
}
