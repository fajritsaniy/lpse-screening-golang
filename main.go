package main

import (
	"fmt"
	"net/http"

	handlers "github.com/fajritsaniy/lpse-screening/handler"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server is starting...")
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/token-generator", handlers.TokenGenerator).Methods("GET")
	r.HandleFunc("/api", handlers.FindProjectAPIHandler).Methods("POST")
	r.HandleFunc("/generate", handlers.GenerateTokenAPIHandler).Methods("POST")

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
