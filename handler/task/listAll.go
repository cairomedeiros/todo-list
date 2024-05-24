package task

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/schemas"
)

func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()

	var tasks []schemas.Task

	result := db.Find(&tasks)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	log.Printf("%d tasks found.\n", result.RowsAffected)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
