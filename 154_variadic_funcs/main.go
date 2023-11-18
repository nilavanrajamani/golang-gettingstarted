package main

import "fmt"

func main() {
	sum := foo([]int{4, 3, 6}...)
	fmt.Println(sum)

	sum = bar([]int{4, 3, 6})
	fmt.Println(sum)
}

func foo(a ...int) int {
	var sum int
	for _, val := range a {
		sum = sum + val
	}
	return sum
}

func bar(a []int) int {
	var sum int
	for _, val := range a {
		sum = sum + val
	}
	return sum
}
