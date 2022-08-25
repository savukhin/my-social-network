package users

import (
	"api/app/utils"
	"api/db/models"
	"api/dto"
	"api/mappers"
	"api/middleware"
	"encoding/json"
	"errors"
	"net/http"
)

func ExtractProfile(req *http.Request) (*models.User, error) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)

	if err != nil {
		return nil, err
	}

	var response *models.User

	if user.ID == 0 {
		user_id := req.Context().Value(middleware.ContextUserIDKey)
		if user_id == nil {
			return nil, errors.New("JWT middleware error")
		}

		response, err = models.GetUserByID(user_id.(int))
	} else {
		response, err = models.GetUserByID(user.ID)
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

func ChangeProfile(res http.ResponseWriter, req *http.Request) {
	changes := &dto.UserEdit{}
	err := json.NewDecoder(req.Body).Decode(changes)

	res.Header().Set("Content-Type", "application/json")

	if err != nil {
		var data = struct {
			Title string `json:"error"`
		}{
			Title: err.Error(),
		}

		res.WriteHeader(400)
		jsonBytes, _ := utils.StructToJSON(data)
		res.Write(jsonBytes)
		return
	}

	id := req.Context().Value(middleware.ContextUserIDKey)
	err = changes.ApllyChanges(id.(int))

	var response interface{}

	if err != nil {
		response = struct {
			Title string "json:\"error\""
		}{Title: err.Error()}

		res.WriteHeader(400)
	} else {
		response = struct {
			Title string "json:\"message\""
		}{Title: "ok"}
	}

	b, _ := json.Marshal(response)
	res.Write(b)
}

func GetProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	current_user_id := req.Context().Value(middleware.ContextUserIDKey)

	response, err := ExtractProfile(req)

	if err != nil {
		var data = struct {
			Title string `json:"error"`
		}{
			Title: err.Error(),
		}

		jsonBytes, _ := utils.StructToJSON(data)
		res.Write(jsonBytes)
		return
	}

	model, _ := response.GetProfile()
	profile := mappers.ToUserProfile(model)

	friendships, err := models.GetFriendships(model.ID)
	if err == nil {

		for _, friendship := range friendships {
			friend_id := friendship.User1ID
			if friendship.User1ID == model.ID {
				friend_id = friendship.User2ID
			}

			if current_user_id != nil && friend_id == current_user_id {
				profile.AddedToFriend = true
			}

			friend_model, err := models.GetUserByID(friend_id)
			if err != nil {
				continue
			}

			compressed := mappers.ToUserCompressed(friend_model)
			profile.Friends = append(profile.Friends, *compressed)
		}

	}

	b, _ := json.Marshal(profile)
	res.Write(b)
}

func AddToFriends(res http.ResponseWriter, req *http.Request) {
	form := &dto.AddToFriends{}
	json.NewDecoder(req.Body).Decode(form)

	user_id := req.Context().Value(middleware.ContextUserIDKey)
	if user_id == nil {
		utils.ResponseError(res, errors.New("no jwt provided"), http.StatusUnauthorized)
		return
	}

	friendship := mappers.AddToFriendsToFriendship(form, user_id.(int))
	_, err := friendship.Save()
	if err != nil {
		utils.ResponseError(res, err, http.StatusInternalServerError)
		return
	}

	utils.ResponseEmptySucess(res)
}

func DeleteFriend(res http.ResponseWriter, req *http.Request) {
	form := &dto.AddToFriends{}
	json.NewDecoder(req.Body).Decode(form)

	user_id := req.Context().Value(middleware.ContextUserIDKey)
	if user_id == nil {
		utils.ResponseError(res, errors.New("no jwt provided"), http.StatusUnauthorized)
		return
	}

	friendship, err := models.GetFriendship(form.AddingUserID, user_id.(int))
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	err = friendship.Delete()
	if err != nil {
		utils.ResponseError(res, err, http.StatusInternalServerError)
		return
	}

	utils.ResponseEmptySucess(res)
}
