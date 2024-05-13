package handler

import "net/http"

func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all ToDo"))
}
