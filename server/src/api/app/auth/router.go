package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", Login).Methods(http.MethodPost)
	router.HandleFunc("/register", Register).Methods(http.MethodPost)
	router.HandleFunc("/check_token", CheckToken).Methods(http.MethodPost)
	return router
}
