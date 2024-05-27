package task

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// @Summary Get task by id
// @Description Get task by id
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id query string true "Task identification"
// @Success 200 {object} handler.CreateTaskRequest{}
// @Failure 400 {object} handler.CreateTaskRequest{}
// @Failure 500 {object} handler.CreateTaskRequest{}
// @Router /task/{id} [get]
func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	task := schemas.Task{}

	if err := db.Preload("SubTasks").First(&task, id).Error; err != nil {
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
