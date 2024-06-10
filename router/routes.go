package router

import (
	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/handler/auth"
	"github.com/cairomedeiros/todo-list/handler/subTask"
	"github.com/cairomedeiros/todo-list/handler/task"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	_ "github.com/cairomedeiros/todo-list/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeRoutes(r *mux.Router) {
	//Initialize Handler
	handler.InitializeHandler()

	//Initialize validator
	validate := validator.New()
	auth := auth.NewAuthHandlerImpl(validate)

	// Protected Routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(auth.AuthMiddleware)

	//tasks
	protected.HandleFunc("/task/create", task.CreateTaskHandler).Methods("POST") //POST
	protected.HandleFunc("/task/listAll", task.ListAllHandler).Methods("GET")    //GET
	protected.HandleFunc("/task/{id}", task.GetByIdHandler).Methods("GET")       //GET
	protected.HandleFunc("/task/{id}", task.UpdateTaskHandler).Methods("PUT")    //PUT
	protected.HandleFunc("/task/{id}", task.DeleteTaskHandler).Methods("DELETE") //DELETE

	//subTask
	protected.HandleFunc("/subTask/create", subTask.CreateSubTaskHandler).Methods("POST") //POST
	protected.HandleFunc("/subTask/{id}", subTask.UpdateSubTaskHandler).Methods("PUT")    //PUT
	protected.HandleFunc("/subTask/{id}", subTask.DeleteSubTaskHandler).Methods("DELETE") //DELETE
	protected.HandleFunc("/subTask/listAll", subTask.ListAllHandler).Methods("GET")       //GET
	protected.HandleFunc("/subTask/{id}", subTask.GetByIdHandler).Methods("GET")          //GET

	//auth
	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")

	protected.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
