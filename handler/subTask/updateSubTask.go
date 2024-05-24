package subTask

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

func UpdateSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	request := handler.UpdateSubTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subTask := schemas.SubTask{}
	if err := db.First(&subTask, id).Error; err != nil {
		http.Error(w, "SubTask not found", http.StatusNotFound)
		return
	}

	if request.TaskID != 0 {
		subTask.TaskID = request.TaskID
	}
	if request.Name != "" {
		subTask.Name = request.Name
	}
	subTask.Completed = request.Completed

	if err := db.Save(&subTask).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "SubTask updated successfully"})

}
