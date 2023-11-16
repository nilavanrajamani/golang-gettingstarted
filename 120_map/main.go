package main

import "fmt"

func main() {
	a := map[string][]string{
		"bond_james":       []string{`shaken, not stirred`, `martinis`, `fast cars`},
		"moneypenny_jenny": []string{`intelligence`, `literature`, `computer science`},
		"no_dr":            []string{`cats`, `ice cream`, `sunsets`},
	}

	for index, value := range a {
		fmt.Println(index)
		for loop_index, loop_value := range value {
			fmt.Println(loop_index, loop_value)
		}
	}

	delete(a, "bond_james")

	fmt.Println("After deletion")

	for index, value := range a {
		fmt.Println(index)
		for loop_index, loop_value := range value {
			fmt.Println(loop_index, loop_value)
		}
	}
}
