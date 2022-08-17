package chat

import (
	"api/app/utils"
	"api/db/models"
	"api/dto"
	"api/mappers"
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

func SendMessage(res http.ResponseWriter, req *http.Request) {
	message := &dto.MessageInput{}
	err := json.NewDecoder(req.Body).Decode(message)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	content, err := mappers.MessageToContent(message)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	_, err = content.Save()
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	utils.ResponseEmptySucess(res)
}

func GetPersonalChatMessages(res http.ResponseWriter, req *http.Request) {
	extraction := &dto.MessageRangeInput{}
	err := json.NewDecoder(req.Body).Decode(extraction)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	messages, err := models.GetMessages(extraction.Offset, extraction.Count, extraction.ChatID)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	messageRange, err := mappers.ToMessageRange(messages)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	b, _ := json.Marshal(messageRange)
	res.Write(b)
}
