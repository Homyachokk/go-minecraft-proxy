package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(client net.Conn) {
	// Установить соединение с сервером Minecraft на localhost:25565
	server, err := net.Dial("tcp", "localhost:25565")
	if err != nil {
		fmt.Println("Не удалось установить соединение с сервером Minecraft:", err)
		return
	}
	defer server.Close()

	// Передать клиентские данные на сервер Minecraft
	go func() {
		_, err := io.Copy(server, client)
		if err != nil {
			fmt.Println("Не удалось передать клиентские данные на сервер Minecraft:", err)
			return
		}
	}()

	// Передать серверные данные клиенту
	_, err = io.Copy(client, server)
	if err != nil {
		fmt.Println("Не удалось передать серверные данные клиенту:", err)
		return
	}
}

func main() {
	// Запустить сервер на localhost:25566
	listener, err := net.Listen("tcp", "localhost:25565")
	if err != nil {
		fmt.Println("Не удалось запустить сервер:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на localhost:25566")

	for {
		// Принять клиентское соединение
		client, err := listener.Accept()
		if err != nil {
			fmt.Println("Не удалось принять клиентское соединение:", err)
			continue
		}

		// Обработать клиентское соединение в отдельной горутине
		go handleConnection(client)
	}
}
