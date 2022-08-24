package posts

import (
	"api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user_posts/{user_id:[0-9]+}", GetUserPosts).Methods(http.MethodGet)

	createPostHandler := http.HandlerFunc(CreateUserPost)
	router.Handle("/create_post", middleware.JwtAuthentication(createPostHandler)).Methods(http.MethodPost)

	likePostHandler := http.HandlerFunc(ToggleLikePost)
	router.Handle("/toggle_like", middleware.JwtAuthentication(likePostHandler)).Methods(http.MethodPost)

	return router
}
