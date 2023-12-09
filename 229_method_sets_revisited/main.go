package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.FirstName, "speaks")
}

func saySomething(human human) {
	human.speak()
}

func main() {
	p := person{FirstName: "firstName1", LastName: "lastName1"}
	saySomething(&p)
}
