package task

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	request := handler.UpdateToDoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if request.Title != "" {
		task.Title = request.Title
	}
	if request.Description != "" {
		task.Description = request.Description
	}
	if request.DueDate != nil {
		task.DueDate = request.DueDate
	}

	if err := db.Save(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ToDo updated successfully"})

}
