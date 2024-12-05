package main

import (
	"healthcare-chatbot/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// Define chatbot endpoint
	r.HandleFunc("/chat", routes.ChatHandler).Methods("POST")

	// Enable CORS
	handler := cors.AllowAll().Handler(r)

	// Start server
	http.ListenAndServe(":8080", handler)
}
