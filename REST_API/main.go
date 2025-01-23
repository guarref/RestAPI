package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type Taskkk struct {
	Juststring string `json:"task"`
}

func HelloHendler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello, ", task)
}

func PostHandler(rw http.ResponseWriter, r *http.Request) {
	var taska Taskkk
	err := json.NewDecoder(r.Body).Decode(&taska)
	if err != nil {
		http.Error(rw, "Неверный json", http.StatusBadRequest)
		return
	}
	task = taska.Juststring
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")
	router.HandleFunc("/api/hello", HelloHendler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
