package posts

import (
	"api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) *mux.Router {
	router.Use(middleware.JwtAuthentication)

	router.HandleFunc("/user_posts/{user_id:[0-9]+}", GetUserPosts).Methods(http.MethodGet)

	return router
}
