package main

import "fmt"

type person struct {
	first string
}

func valueSemantics(a person, s string) person {
	a.first = s
	return a
}

func pointerSemantics(a *person, s string) {
	a.first = s
}

func main() {
	p := person{first: "person1"}
	fmt.Println(valueSemantics(p, "value1").first)
	p2 := person{first: "person2"}
	pointerSemantics(&p2, "value2")
	fmt.Println(p2.first)
	p3 := person{first: "person3"}
	valueSemantics(p3, "value3")
	fmt.Println(p3.first)
}
