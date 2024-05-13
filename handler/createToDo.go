package handler

import "net/http"

func CreateToDoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new ToDO"))
}
