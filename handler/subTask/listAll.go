package subTask

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
)

// @BasePath /api/v1

// @Summary List SubTasks
// @Description List all SubTasks
// @Tags SubTasks
// @Accept json
// @Produce json
// @Success 200 {array} handler.ListSubTasksResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /subTask/listAll [get]
func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	var subTasks []schemas.SubTask

	result := db.Find(&subTasks)

	if result.Error != nil {
		handler.SendError(w, http.StatusInternalServerError, result.Error.Error())
	}

	log.Printf("%d SubTasks found.\n", result.RowsAffected)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(subTasks); err != nil {
		handler.SendError(w, http.StatusBadRequest, err.Error())
	}
}
