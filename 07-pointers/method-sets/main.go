package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		length: 12,
	}

	c := circle{
		radius: 5,
	}

	info(s)
	info(&c)
}
