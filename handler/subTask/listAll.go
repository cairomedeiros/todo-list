package subTask

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
)

func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	var subTasks []schemas.SubTask

	result := db.Find(&subTasks)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	log.Printf("%d SubTasks found.\n", result.RowsAffected)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(subTasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
