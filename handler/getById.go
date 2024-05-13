package handler

import "net/http"

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get by ToDo by id"))
}
