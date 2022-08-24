package posts

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
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserPosts(res http.ResponseWriter, req *http.Request) {
	user_id, err := strconv.Atoi(mux.Vars(req)["user_id"])
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	contents, err := models.GetPostsByUserID(user_id)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	posts := make([]*dto.Post, 0)
	for _, content := range contents {
		post, err := mappers.ContentToPost(content)
		if err == nil {
			posts = append(posts, post)
		}
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(posts)
	res.Write(b)
}

func CreateUserPosts(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	form := &dto.PostCreate{}

	err := json.NewDecoder(req.Body).Decode(form)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	user_id := req.Context().Value(middleware.ContextUserIDKey)
	if user_id == nil {
		utils.ResponseError(res, errors.New(""), http.StatusUnauthorized)
		return
	}

	content, err := mappers.PostCreateToContent(form, user_id.(int))
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	_, err = content.Save()
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	post, err := mappers.ContentToPost(content)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(post)
	res.Write(b)
}
