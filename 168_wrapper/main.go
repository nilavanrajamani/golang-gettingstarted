package main

import (
	"fmt"
	"time"
)

func main() {
	timerFunction(observationFunction)
}

func observationFunction() {
	fmt.Println(time.Second)
	time.Sleep(2 * time.Second)
	fmt.Println("Function completed")
}

func timerFunction(innerfn func()) {
	a := time.Now()
	innerfn()
	elapsed := time.Since(a)
	fmt.Println("The time elapsed for execution is", elapsed)
}
