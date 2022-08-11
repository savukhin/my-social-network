package auth

import (
	"api/app/utils"
	"api/db/models"
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
			Title: "Error",
		}

		jsonBytes, _ := utils.StructToJSON(data)
		res.Write(jsonBytes)
		return
	}
	response := user.Login()
	b, _ := json.Marshal(response)
	res.Write(b)
}

func Register(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		return
	}

	response := user.Register()

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}

func CheckToken(res http.ResponseWriter, req *http.Request) {
	jwt := req.Header.Get("Authorization")
	token, err := utils.UnpackJWT(jwt)
	res.Header().Add("Content-Type", "application/json")
	response := map[string]interface{}{
		"id_token":   "",
		"expires_at": "",
		"message":    "ok",
	}

	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		response["message"] = err
		response["id_token"] = jwt
	} else {
		response["id_token"] = jwt
		response["expires_at"] = token.TimeExp
	}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}
