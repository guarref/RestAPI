package userService

import (
	"pet_project_final/internal/taskService"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Task     []taskService.Task
}
