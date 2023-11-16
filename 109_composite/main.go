package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range a {
		fmt.Printf("i is %v and type is %T\n", v, v)
	}
}
