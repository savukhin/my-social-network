package auth

import (
	"api/app/utils"
	"api/db/models"
	"api/dto"
	"api/mappers"
	"encoding/json"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	res.Header().Set("Content-Type", "application/json")

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
	response, status := user.Login()
	if !status {
		res.WriteHeader(http.StatusBadRequest)
	}
	b, _ := json.Marshal(response)
	res.Write(b)
}

func Register(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	response, status := user.Register()

	res.Header().Add("Content-Type", "application/json")
	if !status {
		res.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(res).Encode(response)
}

func CheckToken(res http.ResponseWriter, req *http.Request) {
	jwt := req.Header.Get("Authorization")
	token, err := utils.UnpackJWT(jwt)
	res.Header().Add("Content-Type", "application/json")

	response := &dto.TokenCheckStruct{}

	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
	} else {
		user, _ := models.GetUserByID(token.UserID)
		response = mappers.ToTokenCheckStruct(user, jwt, token.TimeExp)
	}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}
