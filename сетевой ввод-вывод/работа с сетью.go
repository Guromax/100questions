package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Слушаем на порту 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при создании слушателя:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Сервер запущен и слушает на порту 8080")

	for {
		// Принимаем входящее подключение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии подключения:", err)
			continue
		}

		// Обработка подключения в отдельной горутине
		go handleConnection(conn)
	}
}

// handleConnection обрабатывает отдельное подключение
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Подключился клиент:", conn.RemoteAddr().String())

	// Отправляем сообщение клиенту
	_, err := io.WriteString(conn, "Привет от сервера!n")
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения:", err)
		return
	}

	fmt.Println("Сообщение отправлено клиенту:", conn.RemoteAddr().String())
}
