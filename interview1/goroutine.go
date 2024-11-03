package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("123")

	fmt.Println("looool")
}
