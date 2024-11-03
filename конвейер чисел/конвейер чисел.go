package main

import "fmt"

func main() {

	firstCh := make(chan int)
	secondCh := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			firstCh <- i
		}
		close(firstCh)
	}()

	go func() {
		for i := range firstCh {
			secondCh <- i * i
		}
		close(secondCh)
	}()

	for i := range secondCh {
		fmt.Println(i)
	}
}
