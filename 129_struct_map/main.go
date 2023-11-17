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

	map_person := make(map[string]person)
	map_person[p1.last_name] = p1
	map_person[p2.last_name] = p2
	fmt.Println(map_person)
	for index, value := range map_person[p1.last_name].ice_cream_flavors {
		fmt.Println(index)
		fmt.Println(value)
	}
	for index, value := range map_person[p2.last_name].ice_cream_flavors {
		fmt.Println(index)
		fmt.Println(value)
	}
}
