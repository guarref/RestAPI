package userService

import "pet_project_final/internal/taskService"

type UserService struct {
	urepo UserRepository
	srepo taskService.TaskRepository
}

func NewUserService(urepo UserRepository, srepo taskService.TaskRepository) *UserService {
	return &UserService{urepo: urepo, srepo: srepo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.urepo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.urepo.GetAllUsers()
}

func (s *UserService) GetTasksUserId(user_id uint) ([]taskService.Task, error) {
	return s.srepo.GetTasksUserId(user_id)
}

func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return s.urepo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.urepo.DeleteUserByID(id)
}
