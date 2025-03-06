package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	GetTasksUserId(user_id uint) ([]Task, error)
	// получаем задачи для конкретного user_id
	UpdateTaskByID(id uint64, newtask Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint64) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksUserId(user_id uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", user_id).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint64, newtask Task) (Task, error) {
	var oldtask Task
	result := r.db.First(&oldtask, id)
	if result.Error != nil {
		return Task{}, result.Error
	}

	// Обновляем только указанные поля
	PatchResult := r.db.Model(&oldtask).Select("is_done", "text").Updates(newtask)
	if PatchResult.Error != nil {
		return Task{}, PatchResult.Error
	}
	// Полная перезапись обоих полей
	// oldtask.Task, oldtask.IsDone = newtask.Task, newtask.IsDone
	// r.db.Save(&oldtask)
	return oldtask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint64) error {
	var deltask Task
	result := r.db.First(&deltask, id)
	if result.Error != nil {
		return result.Error
	}
	r.db.Delete(&deltask)
	return nil
}
