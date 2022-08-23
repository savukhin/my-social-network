package chat

import (
	"api/app/utils"
	"api/db/models"
	"api/dto"
	"api/mappers"
	"api/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getLastChatMessage(chat_id int) (*dto.Message, error) {
	messages, err := models.GetMessages(0, 1, chat_id)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, nil
	}

	message_dto, err := mappers.ToMessage(&messages[0])
	if err != nil {
		return nil, err
	}

	return message_dto, nil
}

func GetChats(res http.ResponseWriter, req *http.Request) {
	user_id := req.Context().Value(middleware.ContextUserIDKey)
	res.Header().Set("Content-Type", "application/json")

	if user_id == nil {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	chat_models, err := models.GetChatsByUserID(user_id.(int))

	if err != nil {
		res.WriteHeader(400)
		return
	}

	chats := make([]*dto.Chat, 0)
	for _, model := range chat_models {
		participants_ids, err := models.GetChatParticipants(model.ID)
		if err != nil {
			continue
		}

		participants, err := models.UserIDsToUsers(participants_ids)
		if err != nil {
			continue
		}

		lastMessages, err := getLastChatMessage(model.ID)
		if err != nil || lastMessages == nil {
			continue
		}

		chat, err := mappers.ToChatDTO(model, participants, lastMessages)
		if err != nil {
			continue
		}

		chats = append(chats, chat)
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(chats)
	res.Write(b)
}

func GetChat(res http.ResponseWriter, req *http.Request) {
	chat_id, err := strconv.Atoi(mux.Vars(req)["chat_id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	chat_model, err := models.GetChatByID(chat_id)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	participants_id, err := models.GetChatParticipants(chat_model.ID)

	if err != nil {
		res.WriteHeader(400)
		return
	}

	participant_models := make([]*models.User, 0)
	for _, id := range participants_id {
		participant, err := models.GetUserByID(id)
		if err != nil {
			utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}

		participant_models = append(participant_models, participant)
	}

	chat, err := mappers.ToChatDTO(chat_model, participant_models, nil)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(chat)
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

	chat_model, err := models.GetPersonalChat(user1_id.(int), user2_id)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	participants_id, err := models.GetChatParticipants(chat_model.ID)

	if err != nil {
		res.WriteHeader(400)
		return
	}

	participant_models := make([]*models.User, 0)
	for _, id := range participants_id {
		participant, err := models.GetUserByID(id)
		if err != nil {
			utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}

		participant_models = append(participant_models, participant)
	}

	chat, err := mappers.ToChatDTO(chat_model, participant_models, nil)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
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
	fmt.Println(messages)

	messageRange, err := mappers.ToMessageRange(messages)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}
	fmt.Println(messages)

	b, _ := json.Marshal(messageRange)
	res.Write(b)
}
