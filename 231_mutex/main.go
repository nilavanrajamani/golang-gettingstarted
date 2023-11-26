package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			a := counter
			a++
			counter = a
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("The counter value is ", counter)
}
