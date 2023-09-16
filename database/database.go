package database

import (
	"api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB - глобальная переменная для хранения соединения с базой данных
var DB *gorm.DB

// Connect устанавливает соединение с базой данных
func Connect() error {
	// переменная для хранения ошибки которая может возникнуть при подключении
	var err error
	// устанавливаем соединение с базой данных зщыепкуы с помощью функции Open
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	// проверяем ошибку установки соединения, и если она не равна nil возвращаем ошибку и выходим из функции
	if err != nil {
		return err
	}

	// если соединение успешно установлено,
	// вызываем функцию AutoMigrate на объекте DB, что позволит создать таблицу для структуры Employee
	DB.AutoMigrate(&models.Employee{})

	return nil
}
