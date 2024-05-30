package subTask

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

// @BasePath /api/v1

// @Summary Update SubTask
// @Description Update a SubTask
// @Tags SubTasks
// @Accept json
// @Produce json
// @Param id path string true "SubTask Identification"
// @Param task body handler.UpdateSubTaskRequest true "SubTask data to Update"
// @Success 200 {object} handler.UpdateSubTaskResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /subTask/{id} [put]
func UpdateSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	request := handler.UpdateSubTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handler.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	subTask := schemas.SubTask{}
	if err := db.First(&subTask, id).Error; err != nil {
		handler.SendError(w, http.StatusNotFound, "Task not found")
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
		handler.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccess(w, "update-subtask", subTask)

}
