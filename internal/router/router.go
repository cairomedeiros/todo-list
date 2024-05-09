package router

import (
	"fmt"
	"net/http"
)

func Initialize() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my api!")
	})

	fmt.Println("Server started in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
