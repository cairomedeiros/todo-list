package task

import (
	"fmt"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

// @BasePath /api/v1

// @Summary Delete Task
// @Description Delete Task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task identification"
// @Success 200 {object} handler.DeleteTaskResponse
// @failure 400 {object} handler.ErrorResponse
// @failure 404 {object} handler.ErrorResponse
// @Router /task/{id} [delete]
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		handler.SendError(w, http.StatusBadRequest, handler.ErrParamIsRequired("id", "pathParameter").Error())
		return
	}

	task := schemas.Task{}

	if err := db.Find(&task, id).Error; err != nil {
		handler.SendError(w, http.StatusNotFound, fmt.Sprintf("Task with id: %s not found", id))
		return
	}

	//delete all subtasks by task id before delete task
	subTask := schemas.SubTask{}
	db.Where("task_id = ?", id).Unscoped().Delete(&subTask)

	db.Unscoped().Delete(&task, id)

	handler.SendSuccess(w, "delete-task", task)
}
