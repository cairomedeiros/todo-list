package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cairomedeiros/todo-list/schemas"
)

func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []schemas.Task

	result := db.Find(&tasks)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	log.Printf("%d records found.\n", result.RowsAffected)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
