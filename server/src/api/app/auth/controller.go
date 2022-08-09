package auth

import (
	"api/app/utils"
	"api/db/models"
	"encoding/json"
	"fmt"
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
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Register",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonBytes)
}
