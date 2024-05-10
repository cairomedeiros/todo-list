package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Initialize() {

	// Initialize a new router
	r := mux.NewRouter()

	initializeRoutes(r)

	fmt.Println("Server started in http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
