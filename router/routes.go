package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRoutes(r *mux.Router) {
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/details", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the details page!"))
}
