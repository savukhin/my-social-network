package users

import (
	"api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) *mux.Router {
	router.Use(middleware.JwtAuthentication)
	router.HandleFunc("/profile", GetProfile).Methods(http.MethodPost)
	router.HandleFunc("/change_profile", ChangeProfile).Methods(http.MethodPost)
	return router
}
