package main

import "fmt"

func main() {
	square := func(n int) int {
		return n * n
	}

	printSquare := func(f func(int) int, param int) string {
		return fmt.Sprint("The value of printSquare is ", f(param))
	}

	fmt.Println(printSquare(square, 25))
}
