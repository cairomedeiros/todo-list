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

	db.Find(&todo, id)
	db.Unscoped().Delete(&todo, id)
}
