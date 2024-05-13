package handler

import "net/http"

func DeleteToDoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete this ToDo"))
}
