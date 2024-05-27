package task

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

// @Summary Delete task
// @Description Delete task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id query string true "Task identification"
// @Success 200 {object} handler.CreateTaskRequest{}
// @Failure 400 {object} handler.CreateTaskRequest{}
// @Failure 500 {object} handler.CreateTaskRequest{}
// @Router /task/{id} [delete]
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	task := schemas.Task{}

	if err := db.Find(&task, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//delete all subtasks by task id before delete task
	subTask := schemas.SubTask{}
	db.Where("task_id = ?", id).Unscoped().Delete(&subTask)

	db.Unscoped().Delete(&task, id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}
