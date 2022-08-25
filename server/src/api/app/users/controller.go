package users

import (
	"api/app/utils"
	"api/db/models"
	"api/dto"
	"api/mappers"
	"api/middleware"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

func ChangeAvatar(res http.ResponseWriter, req *http.Request) {
	user_id := req.Context().Value(middleware.ContextUserIDKey)
	if user_id == nil {
		utils.ResponseError(res, errors.New("no jwt provided"), http.StatusUnauthorized)
		return
	}

	user, err := models.GetUserByID(user_id.(int))
	if err != nil {
		utils.ResponseError(res, errors.New("no jwt provided"), http.StatusUnauthorized)
		return
	}

	file, filename, err := req.FormFile("avatar")
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = os.MkdirAll("./uploads/avatars", os.ModePerm)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	path := fmt.Sprintf("./uploads/avatars/%d%s", time.Now().UnixNano(), filepath.Ext(filename.Filename))

	dst, err := os.Create(path)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	content := models.CreateAvatarContent(path, user_id.(int))
	_, err = content.Save()
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	user.Avatar_ID = sql.NullInt64{Int64: int64(content.ID), Valid: true}
	_, err = user.Apply()
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	utils.ResponseEmptySucess(res)
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

func GetAvatar(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	filename, err := models.GetUserAvatarURL(user_id)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	http.ServeFile(res, req, filename)
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
