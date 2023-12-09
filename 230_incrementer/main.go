package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			a := counter
			runtime.Gosched()
			a++
			counter = a
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("The counter value is ", counter)
}
