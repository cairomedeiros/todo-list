package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

func DeleteToDoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	task := schemas.Task{}

	if err := db.Find(&task, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Unscoped().Delete(&task, id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ToDo deleted successfully"})
}
