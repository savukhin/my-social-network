package chat

import (
	"api/app/utils"
	"api/db/models"
	"api/middleware"
	"encoding/json"
	"net/http"
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

	response, err := utils.StructToJSON(chats)
	if err != nil {
		res.WriteHeader(400)
		return
	}

	json.NewEncoder(res).Encode(response)
}
