package main

import (
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"

	"go_final_project/pkg/db"
	"go_final_project/pkg/handlers"
	"go_final_project/pkg/server"
)

func main() {
	// Инициализация базы данных
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных: %v", err)
	}

	// Инициализация маршрутизатора
	r := chi.NewRouter()
	r.Get("/*", handlers.ServeFile)
	r.Get("/", handlers.ServeFile)

	// Запуск сервера
	err = server.Run()
	if err != nil {
		fmt.Print("Ошибка запуска сервера")
	}
}
