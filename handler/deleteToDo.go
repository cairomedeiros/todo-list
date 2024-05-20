package handler

import (
	"net/http"

	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/gorilla/mux"
)

func DeleteToDoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todo := schemas.ToDo{}

	if err := db.Find(&todo, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Unscoped().Delete(&todo, id)
}
