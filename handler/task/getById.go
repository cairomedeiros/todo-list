package task

import (
	"errors"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// @BasePath /api/v1

// @Summary Show Task
// @Description Show a Task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task identification"
// @Success 200 {object} handler.GetTaskByIdResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Router /task/{id} [get]
func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		handler.SendError(w, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	task := schemas.Task{}

	if err := db.Preload("SubTasks").First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.SendError(w, http.StatusNotFound, "Task not found")
			return
		}

		handler.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	handler.SendSuccess(w, "get-task-by-id", task)
}
