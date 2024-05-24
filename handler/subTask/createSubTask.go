package subTask

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
)

func CreateSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	request := handler.CreateSubTaskRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subTask := schemas.SubTask{
		TaskID:    request.TaskID,
		Name:      request.Name,
		Completed: request.Completed,
	}
	if err := db.Create(&subTask).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "SubTask received and saved successfully"})

}
