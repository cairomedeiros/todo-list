package router

import (
	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/handler/subTask"
	"github.com/cairomedeiros/todo-list/handler/task"
	"github.com/gorilla/mux"

	_ "github.com/cairomedeiros/todo-list/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Clean GO API Docs
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:port
// @BasePath /
func initializeRoutes(r *mux.Router) {
	//Initialize Handler
	handler.InitializeHandler()

	//tasks
	r.HandleFunc("/task/create", task.CreateTaskHandler).Methods("POST") //POST
	r.HandleFunc("/task/listAll", task.ListAllHandler).Methods("GET")    //GET
	r.HandleFunc("/task/{id}", task.GetByIdHandler).Methods("GET")       //GET
	r.HandleFunc("/task/{id}", task.UpdateTaskHandler).Methods("PUT")    //PUT
	r.HandleFunc("/task/{id}", task.DeleteTaskHandler).Methods("DELETE") //DELETE

	//subTask
	r.HandleFunc("/subTask/create", subTask.CreateSubTaskHandler).Methods("POST") //POST
	r.HandleFunc("/subTask/{id}", subTask.UpdateSubTaskHandler).Methods("PUT")    //PUT
	r.HandleFunc("/subTask/{id}", subTask.DeleteSubTaskHandler).Methods("DELETE") //DELETE
	r.HandleFunc("/subTask/listAll", subTask.ListAllHandler).Methods("GET")       //GET
	r.HandleFunc("/subTask/{id}", subTask.GetByIdHandler).Methods("GET")          //GET

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
