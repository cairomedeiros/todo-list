package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

func UpdateToDoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	request := UpdateToDoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := schemas.ToDo{}
	if err := db.First(&todo, id).Error; err != nil {
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}

	if request.Title != "" {
		todo.Title = request.Title
	}
	if request.Description != "" {
		todo.Description = request.Description
	}

	if err := db.Save(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ToDo updated successfully"})

}
