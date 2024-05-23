package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	task := schemas.Task{}

	if err := db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
