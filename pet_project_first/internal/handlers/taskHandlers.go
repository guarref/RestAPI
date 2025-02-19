package handlers

import (
	"encoding/json"
	"pet_project_first/internal/taskService" // Импортируем наш сервис
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *taskService.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasksHandler(rw http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(tasks)
}

func (h *Handler) PostTaskHandler(rw http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(createdTask)
}

func (h *Handler) PatchTaskHandler(rw http.ResponseWriter, r *http.Request) {
	var newtask taskService.Task
	err := json.NewDecoder(r.Body).Decode(&newtask)
	if err != nil {
		http.Error(rw, "Неверный json", http.StatusBadRequest)
		return
	}
	// vars := mux.Vars(r)
	// id := vars["id"]
	id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	patchedTask, err := h.Service.UpdateTaskByID(id, newtask)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(patchedTask)
}

func (h *Handler) DeleteTaskHandler(rw http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNoContent)
}
