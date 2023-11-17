package main

import "fmt"

func main() {
	p := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "first1",
		friends:   map[string]int{"key1": 43},
		favDrinks: []string{"drink1", "drink2"},
	}

	fmt.Println(p)
}
