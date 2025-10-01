package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	_, err := s.repo.GetTasksUserId(task.ID)
	if err != nil {
		return Task{}, nil
	}
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint64, newtask Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, newtask)
}

func (s *TaskService) DeleteTaskByID(id uint64) error {
	return s.repo.DeleteTaskByID(id)
}
