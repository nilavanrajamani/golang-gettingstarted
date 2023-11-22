package main

import "fmt"

type youngin interface {
	walk()
	run()
}

type animal struct {
	name string
}

func (p animal) walk() {
	fmt.Println("The animal", p.name, "walks on two legs")
}

func (p animal) run() {
	fmt.Println("The animal", p.name, "runs on four legs")
}

func walk(c youngin) {
	c.walk()
}

func run(c youngin) {
	c.run()
}

func main() {
	p1 := animal{name: "Dog"}
	walk(p1)
	run(p1)

	p2 := &animal{name: "Cat"}
	walk(p2)
	run(p2)
}
