package task

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
)

// @BasePath /api/v1

// @Summary List Tasks
// @Description List all Tasks
// @Tags Tasks
// @Accept json
// @Produce json
// @Success 200 {array} handler.ListTasksResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /task/listAll [get]
func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	var tasks []schemas.Task

	result := db.Find(&tasks)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		handler.SendError(w, http.StatusInternalServerError, result.Error.Error())
	}

	log.Printf("%d tasks found.\n", result.RowsAffected)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		handler.SendError(w, http.StatusBadRequest, err.Error())
	}
}
