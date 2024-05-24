package task

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	request := handler.CreateTaskRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := schemas.Task{
		Title:       request.Title,
		Description: request.Description,
		DueDate:     request.DueDate,
	}
	if err := db.Create(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task received and saved successfully"})

}
