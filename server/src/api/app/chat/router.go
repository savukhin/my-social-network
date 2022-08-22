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
	router.HandleFunc("/chat/getMessages", GetPersonalChatMessages).Methods(http.MethodPost)
	router.HandleFunc("/chat/sendMessage", SendMessage).Methods(http.MethodPost)

	// hub := CreateHub()
	// go hub.Run()

	// router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	// router.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Got websocked")
	// 	ServeWs(hub, w, r)
	// })

	// router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	// an example API handler
	// 	err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// })
	return router
}
