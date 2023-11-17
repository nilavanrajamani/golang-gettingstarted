package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  string
	color  string
}

func main() {
	engine := engine{electric: true}
	v1 := vehicle{engine: engine, make: "make1", model: "model1", doors: "doors1", color: "colors1"}

	fmt.Println(v1)
	fmt.Println(v1.engine, v1.engine.electric, v1.make, v1.model, v1.doors, v1.color)
}
