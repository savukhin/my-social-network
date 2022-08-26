package users

import (
	"api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/get_avatar/{user_id:[0-9]+}", GetAvatar).Methods(http.MethodGet)
	router.HandleFunc("/get_friends/{user_id:[0-9]+}", GetFriends).Methods(http.MethodGet)

	r := router.PathPrefix("").Subrouter()
	r.Use(middleware.JwtAuthentication)
	r.HandleFunc("/profile", GetProfile).Methods(http.MethodPost)
	r.HandleFunc("/change_profile", ChangeProfile).Methods(http.MethodPost)
	r.HandleFunc("/change_avatar", ChangeAvatar).Methods(http.MethodPost)
	r.HandleFunc("/add_to_friend", AddToFriends).Methods(http.MethodPost)
	r.HandleFunc("/delete_friend", DeleteFriend).Methods(http.MethodPost)

	return router
}
