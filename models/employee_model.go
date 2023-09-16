package models

// Employee структура для хранения данных сотрудника
type Employee struct {
	EmployeeID   uint   `gorm:"primary_key" json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	EmployeeAge  uint   `json:"employee_age"`
}
