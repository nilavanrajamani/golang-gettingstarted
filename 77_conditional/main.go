package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	switch {
	case x > 0 && x < 100:
		fmt.Println("The value is between 0 and 100")
	case x > 101 && x < 200:
		fmt.Println("The value is between 101 and 200")
	case x > 201 && x < 250:
		fmt.Println("The value is between 201 and 250")
	default:
		fmt.Println("The value is greater than 250")
	}
}
