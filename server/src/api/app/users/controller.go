package users

import (
	"api/app/utils"
	"api/db/models"
	"api/dto"
	"api/mappers"
	"api/middleware"
	"encoding/json"
	"errors"
	"fmt"
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
	fmt.Println("Change profile")

	if err != nil {
		var data = struct {
			Title string `json:"error"`
		}{
			Title: "Error",
		}

		jsonBytes, _ := utils.StructToJSON(data)
		res.Write(jsonBytes)
		return
	}

	id := req.Context().Value(middleware.ContextUserIDKey)
	err = changes.ApllyChanges(id.(int))

	var response interface{}

	if err != nil {
		fmt.Println(err)
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
	// var response interface{}
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

	profile, _ := response.GetProfile()
	b, _ := json.Marshal(mappers.ToUserProfile(profile))
	res.Write(b)
}
