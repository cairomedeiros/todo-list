package handler

import "net/http"

func UpdateToDoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update ToDo"))
}
