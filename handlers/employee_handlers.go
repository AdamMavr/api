package handlers

import (
	"api/database"
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetEmployee - обработчик GET-запроса для получения сотрудника по ID
func GetEmployee(c *gin.Context) {
	id := c.Param("id") // извлекаем значение параметра "id" из url-пути
	// создаем переменную структуры Employee. В ней будут храниться данные найденного сотрудника
	var employee models.Employee

	// ищем сотрудника в базе данных по id. Метод First возвращает первую запись удовлетворяющую условию
	if err := database.DB.First(&employee, id).Error; err != nil {
		// если сотрудник не найден, возвращаем http-статус 404 и json-ответ с ошибкой
		c.JSON(http.StatusNotFound, gin.H{"error": "Такого сотрудника не существует"})
		return // выходим из функции
	}
	c.JSON(http.StatusOK, employee) // если сотрудник найден - возвращаем http-статус 200 и данные сотрудника в формате json
}

// PostEmployee - обработчик POST-запроса для создания сотрудника
func PostEmployee(c *gin.Context) {
	// создаем переменную структуры Employee. В ней будут храниться данные созданного сотрудника
	var employee models.Employee
	// с помощью c.ShouldBindJSON(&employee) проверяем корректность json-запроса и привязке данных к переменной employee
	// также проверяем ошибку возвращаемого значения ShouldBindJSON. Если ошибка не равна nil, значит запрос некорректный
	// и в таком случае мы возвращаем http-статус 400 и json-ответ с ошибкой
	if err := c.ShouldBindJSON(&employee); err != nil {
		// если запрос некорректный, возвращаем http-статус 400 и json-ответ с ошибкой
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
		return // выходим из функции
	}
	// создаем нового сотрудника в базе данных
	// метод Create сохраняет данные из переменной employee в базу данных
	// также проверяем ошибку возвращаемого Create значения
	// если она не равна nil, значит сотрудника не удалось создать
	// в таком случае возвращаем http-статус 500 и json-ответ с ошибкой
	if err := database.DB.Create(&employee).Error; err != nil {
		// если не удалось создать сотрудника - возвращаем http-статус 500 и json-ответ с ошибкой
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать сотрудника"})
		return // выходим из функции
	}
	// если сотрудника удалось создать - возвращаем json-ответ с http-статусом 201 и данными нового сотрудника в формате json
	c.JSON(http.StatusCreated, employee)
}
