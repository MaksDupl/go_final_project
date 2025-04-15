package main

import (
	"fmt"
	"log"

	"go_final_project/pkg/db"
	"go_final_project/pkg/server"
)

func main() {
	// Инициализация базы данных
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных: %v", err)
	}

	// Запуск сервера
	err = server.Run()
	if err != nil {
		fmt.Print("Ошибка запуска сервера")
	}
}
