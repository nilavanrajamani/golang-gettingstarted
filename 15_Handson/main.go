package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

var a = "variable1"

const b = "variable2"

func main() {
	c := "variable3"
	puppy.Bark()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
