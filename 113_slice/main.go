package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"MoneyPenny", "Im 008"}

	z := [][]string{x, y}

	for _, i := range z {
		fmt.Println(i)
	}

	fmt.Println(z)
}
