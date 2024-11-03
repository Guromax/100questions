package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomCounter struct {
	Val int64
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.Val)
}

func main() {
	X := 100
	Y := 4
	var waitGroup sync.WaitGroup
	counter := atomCounter{}

	for i := 0; i < X; i++ {
		waitGroup.Add(1)
		go func(no int) {
			defer waitGroup.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.Val, 1)
			}
		}(i)
		waitGroup.Wait()
		fmt.Println(counter.Value())
	}

}
