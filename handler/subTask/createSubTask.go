package subTask

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
)

// @BasePath /api/v1

// @Summary Create SubTask
// @Description Create a new SubTask
// @Tags SubTasks
// @Accept json
// @Produce json
// @Param request body handler.CreateSubTaskRequest false "Request body"
// @Success 200 {object} handler.SubTaskResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /subTask/create [post]
func CreateSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	request := handler.CreateSubTaskRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handler.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	subTask := schemas.SubTask{
		TaskID:    request.TaskID,
		Name:      request.Name,
		Completed: request.Completed,
	}
	if err := db.Create(&subTask).Error; err != nil {
		handler.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "SubTask received and saved successfully"})

}
