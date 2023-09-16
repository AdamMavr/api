package handlers

import (
	"api/database"
	"api/models"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// GetEmployee - обработчик GET-запроса для получения сотрудника по ID
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Такого работника не существует"})
		return
	}
	serializedEmployee, err := json.Marshal(employee)
	if err != nil {
		log.Fatalf("Конвертация в JSON не удалась: %v", err)
	}
	c.Data(http.StatusOK, "application/json", serializedEmployee)
}

// PostEmployee - обработчик POST-запроса для создания сотрудника
func PostEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
		return
	}
	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать работника"})
		return
	}
	serializedEmployee, err := json.Marshal(employee)
	if err != nil {
		log.Fatalf("Конвертация в JSON не удалась: %v", err)
	}
	c.Data(http.StatusCreated, "application/json", serializedEmployee)
}
