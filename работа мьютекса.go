package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			fmt.Println(counter)
		}()
	}
}
