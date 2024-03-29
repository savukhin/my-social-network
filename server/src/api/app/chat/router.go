package chat

import (
	"api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) *mux.Router {
	router.Use(middleware.JwtAuthentication)
	router.HandleFunc("/chats", GetChats).Methods(http.MethodGet)
	router.HandleFunc("/chat/by_user/{user_id:[0-9]+}", GetPersonalChat).Methods(http.MethodPost)
	router.HandleFunc("/chat/by_id/{chat_id:[0-9]+}", GetChat).Methods(http.MethodPost)
	router.HandleFunc("/chat/getMessages", GetPersonalChatMessages).Methods(http.MethodPost)
	router.HandleFunc("/chat/sendMessage", SendMessage).Methods(http.MethodPost)

	return router
}
