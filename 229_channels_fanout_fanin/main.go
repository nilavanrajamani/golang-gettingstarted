package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func(k int) {
				k = k * 10
				for j := k; j < (k + 10); j++ {
					a <- j
				}
				wg.Done()
			}(i)
		}

		wg.Wait()
		close(a)
	}()

	for b := range a {
		fmt.Println(b)
	}

	fmt.Println("About to exit")
}
