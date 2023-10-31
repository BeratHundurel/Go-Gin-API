package main

import (
	"example/data-acces/api/routes"
	"example/data-acces/internal/database"
	"log"
	"net/http"
)

func main() {
	router := routes.SetRoutes()
	database.InitDB()
	// Start the web server
	port := ":8080" // Define the port you want to listen on
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
