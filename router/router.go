package router

import (
	"authservice/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// add user endpoints
	router.HandleFunc("/v1/signIn", controller.SignIn).Methods("POST")
	router.HandleFunc("/v1/validateToken", controller.ValidateToken).Methods("POST")
	router.HandleFunc("/v1/refreshToken", controller.RefreshToken).Methods("POST")

	return router
}
