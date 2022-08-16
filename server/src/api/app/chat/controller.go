package chat

import (
	"api/db/models"
	"api/middleware"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetChats(res http.ResponseWriter, req *http.Request) {
	user_id := req.Context().Value(middleware.ContextUserIDKey)
	res.Header().Set("Content-Type", "application/json")

	if user_id == nil {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	chats, err := models.GetChatsByUserID(user_id.(int))

	if err != nil {
		res.WriteHeader(400)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(chats)
	res.Write(b)
}

func GetPersonalChat(res http.ResponseWriter, req *http.Request) {
	user1_id := req.Context().Value(middleware.ContextUserIDKey)
	if user1_id == nil {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	user2_id, err := strconv.Atoi(mux.Vars(req)["user_id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	chat, err := models.GetPersonalChat(user1_id.(int), user2_id)

	if err != nil {
		res.WriteHeader(400)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(chat)
	res.Write(b)
}
