package subTask

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// @BasePath /api/v1

// @Summary Show SubTask
// @Description Show a SubTask
// @Tags SubTasks
// @Accept json
// @Produce json
// @Param id path string true "SubTask identification"
// @Success 200 {object} handler.GetSubTaskByIdResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Router /subTask/{id} [get]
func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		handler.SendError(w, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	subTask := schemas.SubTask{}

	if err := db.First(&subTask, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.SendError(w, http.StatusNotFound, "SubTask not found")
			return
		}

		handler.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subTask)
}
