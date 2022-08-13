package users

import (
	"api/app/utils"
	"api/db/models"
	"encoding/json"
	"net/http"
)

func ChangeProfile(res http.ResponseWriter, req *http.Request) {
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
	response := user.GetProfile()
	b, _ := json.Marshal(response)
	res.Write(b)
}

func GetProfile(res http.ResponseWriter, req *http.Request) {
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
	response := user.GetProfile()
	b, _ := json.Marshal(response)
	res.Write(b)
}
