package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := truck{
		vehicle: vehicle{
			doors: 6,
			color: "red",
		},
		fourWheel: true,
	}

	v2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "grey",
		},
		luxury: true,
	}

	fmt.Println(v1)
	fmt.Println(v2)

}
