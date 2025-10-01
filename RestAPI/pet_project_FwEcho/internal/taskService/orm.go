package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Text   string `json:"text"`    // Наш сервер будет ожидать json c полем text
	IsDone bool   `json:"is_done"` // В GO используем CamelCase, в Json - snake
}
