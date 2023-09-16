package main

import (
	"api/database"
	"api/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// подключение к базе данных
	// вызываем функцию Connect из пакета database
	// если во время подулючения произошла ошибка, выполнение программы прерывается
	// и выводится ошибка через log
	err := database.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// создание маршрутизатора Gin
	// экземпляр маршрутизатора по умолчанию
	router := gin.Default()

	// Определение обработчиков запросов
	// GET-маршрут для получения данных о стотруднике из таблицы employees по id
	router.GET("/employees/:id", handlers.GetEmployee)
	// POST-маршрут для создания нового сотрудника для таблицы employees
	router.POST("/employees", handlers.PostEmployee)

	// Запуск сервера на порту 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
