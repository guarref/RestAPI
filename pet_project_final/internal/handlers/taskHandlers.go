package handlers

import (
	"context"
	"pet_project_final/internal/taskService" // Импортируем наш сервис
	"pet_project_final/internal/userService"
	"pet_project_final/internal/web/tasks"
)

type Handler struct {
	Service  *taskService.TaskService
	UService *userService.UserService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService, userservice *userService.UserService) *Handler {
	return &Handler{
		Service:  service,
		UService: userservice,
	}
}

// GetTasksUserId implements tasks.StrictServerInterface.
func (h *Handler) GetTasksUserId(ctx context.Context, request tasks.GetTasksUserIdRequestObject) (tasks.GetTasksUserIdResponseObject, error) {

	alltasks, err := h.UService.GetTasksUserId(request.UserId)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksUserId200JSONResponse{}

	for _, tsk := range alltasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Text:   &tsk.Text,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Text:   &tsk.Text,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Text:   *taskRequest.Text,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Text:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {

	id := uint64(request.Id)

	taskToUpdate := taskService.Task{
		Text:   *request.Body.Text,
		IsDone: *request.Body.IsDone,
		UserID: *request.Body.UserId,
	}
	// taskToUpdate := taskService.Task{}

	// if request.Body.Text != nil {
	// 	taskToUpdate.Text = *request.Body.Text
	// }
	// if request.Body.IsDone != nil { // 🔥 Добавили проверку
	// 	taskToUpdate.IsDone = *request.Body.IsDone
	// }

	patchedTask, err := h.Service.UpdateTaskByID(id, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &patchedTask.ID,
		Text:   &patchedTask.Text,
		IsDone: &patchedTask.IsDone,
		UserId: &patchedTask.UserID,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {

	id := uint64(request.Id)

	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}
