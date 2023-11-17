package main

import "fmt"

type person struct {
	first_name        string
	last_name         string
	ice_cream_flavors []string
}

func main() {
	p1 := person{first_name: "FirstName1", last_name: "LastName1", ice_cream_flavors: []string{"Vanilla", "Chocolate"}}
	p2 := person{first_name: "FirstName2", last_name: "LastName2", ice_cream_flavors: []string{"BlueBerry", "Almond"}}

	fmt.Println(p1)
	fmt.Println(p2)
	for index, value := range p1.ice_cream_flavors {
		fmt.Println(index)
		fmt.Println(value)
	}
	for index, value := range p2.ice_cream_flavors {
		fmt.Println(index)
		fmt.Println(value)
	}
}
