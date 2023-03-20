package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Luke","Last":"Skywalker","Age":32},{"First":"Anakin","Last":"Skywalker","Age":60}]`
	bs := []byte(s)

	//people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Println("Person Number -> ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
