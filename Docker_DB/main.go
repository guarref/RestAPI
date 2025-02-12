package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	json.NewEncoder(rw).Encode(taska) // обратно кодируем в json и отправляем клиенту
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB() // подключаемся к БД

	// Автоматическая миграция модели Task (создание таблицы в БД)
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter() // создаем роутер для обработки запросов
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}
