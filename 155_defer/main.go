package main

import "fmt"

func main() {
	fmt.Println("Executing function 1 ")
	defer exit(1)
	fmt.Println("Executing function 2 ")
	defer exit(2)
	fmt.Println("Executing function 3 ")
	defer exit(3)
}

func exit(assign int) {
	fmt.Println("exit of function", assign)
}
