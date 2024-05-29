package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/rajivgeraev/routes"
)

func main() {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	// Инициализация маршрутов
	routes.InitializeRoutes(router)

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8080", router))
}
