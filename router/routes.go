package router

import (
	"github.com/cairomedeiros/todo-list/handler"
	"github.com/gorilla/mux"
)

func initializeRoutes(r *mux.Router) {
	//Initialize Handler
	handler.InitializeHandler()

	r.HandleFunc("/createTask", handler.CreateToDoHandler).Methods("POST") //POST
	r.HandleFunc("/listAll", handler.ListAllHandler).Methods("GET")        //GET
	r.HandleFunc("/{id}", handler.GetByIdHandler).Methods("GET")           //GET
	r.HandleFunc("/{id}", handler.UpdateToDoHandler).Methods("PUT")        //PUT
	r.HandleFunc("/{id}", handler.DeleteToDoHandler).Methods("DELETE")     //DELETE
}
