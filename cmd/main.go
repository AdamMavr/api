package main

import (
	"log"

	"api/database"
	"api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к базе данных
	err := database.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Создание маршрутизатора Gin
	router := gin.Default()

	// Определение обработчиков запросов
	router.GET("/employees/:id", handlers.GetEmployee)
	router.POST("/employees", handlers.PostEmployee)

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
