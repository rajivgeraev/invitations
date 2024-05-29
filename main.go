package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/rajivgeraev/invitations/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	routes.InitializeRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
