package main

import (
	"fmt"
	"sync"
	"time"
)

var cache = make(map[string]string)
var mu sync.RWMutex

func set(key string, value string) {
	mu.Lock()
	cache[key] = value
	mu.Unlock()
}

func get(key string) string {
	mu.RLock()
	defer mu.RUnlock()
	return cache[key]
}

func main() {
	set("name", "John")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(get("name"))
			wg.Done()
		}()
	}

	time.Sleep(1 * time.Second)
	set("name", "Doe")

	wg.Wait()
}
