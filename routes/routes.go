package routes

import (
	"github.com/gorilla/mux"
	"github.com/rajivgeraev/controllers"
)

func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/invitations/{code}", controllers.RegisterInvitation).Methods("POST")
}
