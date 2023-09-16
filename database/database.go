package database

import (
	"api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// Connect устанавливает соединение с базой данных
func Connect() error {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		return err
	}

	// Миграция таблицы employees
	DB.AutoMigrate(&models.Employee{})

	return nil
}
