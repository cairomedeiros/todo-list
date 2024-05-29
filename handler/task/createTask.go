package task

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
)

// @BasePath /api/v1

// @Summary Create Task
// @Description Create a new Task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param request body handler.CreateTaskRequest false "Request body"
// @Success 200 {object} schemas.TaskResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /task/create [post]
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	request := handler.CreateTaskRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handler.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	task := schemas.Task{
		Title:       request.Title,
		Description: request.Description,
		DueDate:     request.DueDate,
		Completed:   request.Completed,
	}
	if err := db.Create(&task).Error; err != nil {
		handler.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task received and saved successfully"})

}
