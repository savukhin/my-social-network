package auth

import (
	"api/app/utils"
	"fmt"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("new request")
	// setupCorsResponse(&w, r)
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Login",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonBytes)
}

func Register(res http.ResponseWriter, req *http.Request) {
	fmt.Println("new request")
	// setupCorsResponse(&w, r)
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
