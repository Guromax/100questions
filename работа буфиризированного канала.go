package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int, 5)
	// канал numbers не может хранить более пяти целых чисел — это буферный канал с емкостью 5
	counter := 10
	for i := 0; i < counter; i++ {
		select {
		// здесь происходит обработка
		case numbers <- i * i:
			fmt.Println("About to process", i)
		default:
			fmt.Print("No space for ", i, " ")
		}
		// мы начинаем помещать данные в numbers, однако когда канал заполнен, он перестанет сохранять данные и будет выполняться ветка default
	}
	fmt.Println()
	for {
		select {
		case num := <-numbers:
			fmt.Print("*", num, " ")
		default:
			fmt.Println("Nothing left to read!")
			return
		}
	}
}
