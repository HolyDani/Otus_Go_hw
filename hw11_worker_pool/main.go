package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %d: Start\n", i)
			mu.Lock()
			counter++
			defer mu.Unlock()
			fmt.Printf("goroutine %d: Finish\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished!")
	fmt.Printf("Counter = %d\n", counter)
}
