package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/schemas"
)

func CreateToDoHandler(w http.ResponseWriter, r *http.Request) {
	request := CreateToDoRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := schemas.ToDo{
		Title:       request.Title,
		Description: request.Description,
	}
	if err := db.Create(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data received and saved successfully"})

}
