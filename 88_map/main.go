package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"MoneyPenny": 32,
	}

	for i, val := range m {
		fmt.Println("Key is ", i)
		fmt.Println("Value is ", val)
	}
}
