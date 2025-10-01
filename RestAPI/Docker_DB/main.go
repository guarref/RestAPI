package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMessages(rw http.ResponseWriter, r *http.Request) {
	var tasks []Task
	result := DB.Find(&tasks)
	if result.Error != nil {
		fmt.Printf("Записи не найдены %v", result.Error)
		return
	}
	json.NewEncoder(rw).Encode(tasks)
}

func CreateMessage(rw http.ResponseWriter, r *http.Request) {
	var taska Task
	err := json.NewDecoder(r.Body).Decode(&taska)
	if err != nil {
		http.Error(rw, "Неверный json", http.StatusBadRequest)
		return
	}
	DB.Create(&taska)                 // отправили в БД структуру Task
	json.NewEncoder(rw).Encode(taska) // обратно кодируем в json и отправляем клиенту, чтобы показать, что все записалось корректно
}

func PatchMessage(rw http.ResponseWriter, r *http.Request) {
	var newtask Task
	err := json.NewDecoder(r.Body).Decode(&newtask)
	if err != nil {
		http.Error(rw, "Неверный json", http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var oldtask Task
	result := DB.First(&oldtask, id)
	if result.Error != nil {
		http.Error(rw, "Нет такой таски", http.StatusNotFound)
		return
	}
	oldtask.Task, oldtask.IsDone = newtask.Task, newtask.IsDone
	DB.Save(&oldtask)
	json.NewEncoder(rw).Encode(oldtask)
}

func DeleteMessage(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var deltask Task
	result := DB.First(&deltask, id)
	if result.Error != nil {
		http.Error(rw, "Нет такой таски", http.StatusNotFound)
		return
	}
	DB.Delete(&deltask)
	rw.WriteHeader(http.StatusNoContent)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB() // подключаемся к БД

	// Автоматическая миграция модели Task (создание таблицы в БД)
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter() // создаем роутер для обработки запросов
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages/{id}", PatchMessage).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteMessage).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
