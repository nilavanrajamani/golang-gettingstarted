package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("BEGIN CPU", runtime.NumCPU())
	fmt.Println("BEGIN SUBROUTINE", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("go func1")
		wg.Done()
	}()

	go func() {
		fmt.Println("go func2")
		wg.Done()
	}()

	fmt.Println("MID SUBROUTINE", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("END CPU", runtime.NumCPU())
	fmt.Println("END SUBROUTINE", runtime.NumGoroutine())
}
