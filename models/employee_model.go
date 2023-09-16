package models

// Employee структура для хранения данных сотрудника
type Employee struct {
	// gorm:primary_key - говорит о том, что id это первичный идентификатор в базе данных
	EmployeeID   uint   `gorm:"primary_key" json:"employee_id"` // при сереализации структуры в json это поле будет называться employee_id
	EmployeeName string `json:"employee_name"`                  // при сереализации структуры в json это поле будет называться employee_name
	EmployeeAge  uint   `json:"employee_age"`                   // при сереализации структуры в json это поле будет называться employee_age
}
