package task

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

// @BasePath /api/v1

// @Summary Update Task
// @Description Update a Task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task Identification"
// @Param task body handler.UpdateTaskRequest true "Task data to Update"
// @Success 200 {object} handler.UpdateTaskResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /task/{id} [put]
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	request := handler.UpdateTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handler.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		handler.SendError(w, http.StatusNotFound, "Task not found")
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
	task.Completed = request.Completed

	if err := db.Save(&task).Error; err != nil {
		handler.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})

}
