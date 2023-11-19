package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p1 person) speak() {
	fmt.Printf("Name of person is %v and the age of the person is %v", p1.first, p1.age)
}

func main() {
	v1 := person{first: "firstname1", age: 12}
	v1.speak()
}
