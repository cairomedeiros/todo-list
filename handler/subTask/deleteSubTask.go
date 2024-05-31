package subTask

import (
	"fmt"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

// @BasePath /api/v1

// @Summary Delete SubTask
// @Description Delete SubTask
// @Tags SubTasks
// @Accept json
// @Produce json
// @Param id path string true "SubTask identification"
// @Success 200 {object} handler.DeleteSubTaskResponse
// @failure 400 {object} handler.ErrorResponse
// @failure 404 {object} handler.ErrorResponse
// @Router /subTask/{id} [delete]
func DeleteSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		handler.SendError(w, http.StatusBadRequest, handler.ErrParamIsRequired("id", "pathParameter").Error())
		return
	}

	subTask := schemas.SubTask{}

	if err := db.Find(&subTask, id).Error; err != nil {
		handler.SendError(w, http.StatusNotFound, fmt.Sprintf("Task with id: %s not found", id))
		return
	}
	db.Unscoped().Delete(&subTask, id)

	handler.SendSuccess(w, "delete-subtask", subTask)
}
