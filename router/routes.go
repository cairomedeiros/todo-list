package router

import (
	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/handler/subTask"
	"github.com/cairomedeiros/todo-list/handler/task"
	"github.com/gorilla/mux"
)

func initializeRoutes(r *mux.Router) {
	//Initialize Handler
	handler.InitializeHandler()

	//tasks
	r.HandleFunc("/createTask", task.CreateTaskHandler).Methods("POST") //POST
	r.HandleFunc("/listAll", task.ListAllHandler).Methods("GET")        //GET
	r.HandleFunc("/{id}", task.GetByIdHandler).Methods("GET")           //GET
	r.HandleFunc("/{id}", task.UpdateTaskHandler).Methods("PUT")        //PUT
	r.HandleFunc("/{id}", task.DeleteTaskHandler).Methods("DELETE")     //DELETE

	//subTask
	r.HandleFunc("/createSubTask", subTask.CreateSubTaskHandler).Methods("POST") //POST
	r.HandleFunc("/{id}", subTask.UpdateSubTaskHandler).Methods("PUT")           //PUT
	r.HandleFunc("/{id}", subTask.DeleteSubTaskHandler).Methods("DELETE")        //DELETE
	r.HandleFunc("/listAll", subTask.ListAllHandler).Methods("GET")              //GET
	r.HandleFunc("/{id}", subTask.GetByIdHandler).Methods("GET")                 //GET
}
