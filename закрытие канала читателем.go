package main

import "fmt"

func main() {
	dataCh := make(chan int)
	stopCh := make(chan struct{})

	go func() {
		for {
			select {
			case data, ok := <-dataCh:
				if !ok {
					// Канал закрыт, прекращаем обработку
					return
				}
				// Обработка данных
				fmt.Println(data)
			case <-stopCh:
				// Получен сигнал остановки, закрываем канал dataCh
				close(dataCh)
				return
			}
		}
	}()

	// Отправка данных в канал
	dataCh <- 1
	dataCh <- 2

	// Отправка сигнала остановки
	stopCh <- struct{}{}
}
