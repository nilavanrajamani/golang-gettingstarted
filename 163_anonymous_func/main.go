package main

import "fmt"

func main() {
	var a = func() func() {
		b := 54
		c := func() {
			b = b + 1
			fmt.Println("The value of b is ", b)
		}
		return c
	}

	var scopedFunction = a()
	scopedFunction()
	scopedFunction()
	scopedFunction()
}
