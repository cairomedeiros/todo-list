package subTask

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

func DeleteSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]

	subTask := schemas.SubTask{}

	if err := db.Find(&subTask, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Unscoped().Delete(&subTask, id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "SubTask deleted successfully"})
}
