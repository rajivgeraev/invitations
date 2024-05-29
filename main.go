package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rajivgeraev/invitations/config"
	"github.com/rajivgeraev/invitations/routes"
)

func main() {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("MONGO_URL is not set in environment variables")
	}

	config.ConnectDB()

	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
