package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	d := []int{56, 57, 58, 59}

	b := append(a, 52)
	fmt.Println(b)

	c := append(b, 53, 54, 55)
	fmt.Println(c)

	fmt.Println(append(c, d...))
}
