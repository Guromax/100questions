package main

import "fmt"

func fibonacci(num uint) uint {
	if num < 2 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	fmt.Println(fibonacci(10))
}
