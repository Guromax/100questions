package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Начало функции main()")
	var wg sync.WaitGroup
	for i := 1; i < 4; i++ {
		wg.Add(1) // Увеличиваем счетчик потоков на единицу
		go func(n int) {
			defer wg.Done() // Уменьшаем счетчик потоков на единицу
			for j := 1; j < 11; j++ {
				fmt.Println("Поток:", n, "j =", j)
				time.Sleep(time.Second) // Имитация выполнения задачи
			}
		}(i)
	}
	wg.Wait() // Ожидаем завершения всех потоков
	fmt.Println("Конец функции main()")
}
