package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"MoneyPenny": 32,
	}

	m1 := m["James"]
	fmt.Println(m1)

	m1 = m["Q"]
	fmt.Println(m1)

	if _, ok := m["Q"]; !ok {
		fmt.Println("No Q Present")
	}
}
