package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func increment() {
	mu.Lock()
	count++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func() {
			increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
