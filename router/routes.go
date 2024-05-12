package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRoutes(r *mux.Router) {
	r.HandleFunc("/createtodo", CreateToDoHandler) //POST
	r.HandleFunc("/listAll", ListAllHandler)       //GET
	r.HandleFunc("/{id}", GetByIdHandler)          //GET
	r.HandleFunc("/{id}", UpdateToDoHandler)       //PUT
	r.HandleFunc("/{id}", DeleteToDoHandler)       //DELETE
}

func CreateToDoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new ToDO"))
}

func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all ToDo"))
}

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get by ToDo by id"))
}

func UpdateToDoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update ToDo"))
}

func DeleteToDoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete this ToDo"))
}
