package users

import (
	"api/app/utils"
	"api/db/models"
	"api/middleware"
	"encoding/json"
	"errors"
	"net/http"
)

func ExtractProfile(req *http.Request) (interface{}, error) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)

	if err != nil {
		return nil, err
	}

	var response interface{}

	if user.ID == 0 {
		user_id := req.Context().Value(middleware.ContextUserIDKey)
		if user_id == nil {
			return nil, errors.New("JWT middleware error")
		}
		response, err = models.GetUserByID(user_id.(int))

		if err != nil {
			return nil, err
		}
	} else {
		response = user.GetProfile()
	}

	return response, nil
}

func ChangeProfile(res http.ResponseWriter, req *http.Request) {
	response, err := ExtractProfile(req)
	res.Header().Set("Content-Type", "application/json")

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

	b, _ := json.Marshal(response)
	res.Write(b)
}

func GetProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	response, err := ExtractProfile(req)

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

	b, _ := json.Marshal(response)
	res.Write(b)
}
