package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{First: "first1", Last: "last1", Age: 5}
	p2 := person{First: "first2", Last: "last2", Age: 5}
	p := []person{p1, p2}
	fmt.Println(p)

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	people := []person{}

	json.Unmarshal(bs, &people)
	fmt.Println(people)

}
