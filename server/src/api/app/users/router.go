package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/profile", GetProfile).Methods(http.MethodPost)
	return router
}
