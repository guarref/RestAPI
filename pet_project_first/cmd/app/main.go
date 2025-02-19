package main

import (
	"net/http"
	"pet_project_first/internal/database"
	"pet_project_first/internal/handlers"
	"pet_project_first/internal/taskService"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/messages", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/messages/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
