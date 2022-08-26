package posts

import (
	"api/app/utils"
	"api/db/models"
	"api/dto"
	"api/mappers"
	"api/middleware"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserPosts(res http.ResponseWriter, req *http.Request) {
	jwt := req.Header.Get("Authorization")
	token, err := utils.UnpackJWT(jwt)
	current_user_id := -1
	if err == nil {
		current_user_id = token.UserID
	}

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
		if err != nil {
			continue
		}

		likes, err := models.GetLikesByContent(post.ID)
		if err != nil {
			posts = append(posts, post)
			continue
		}

		for _, model := range likes {
			if like, err := mappers.ToLikePost(model); err == nil {
				post.Likes = append(post.Likes, like)

				if like.UserID == current_user_id {
					post.HasCurrentUserLike = true
				}
			}
		}

		posts = append(posts, post)
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(posts)
	res.Write(b)
}

func CreateUserPost(res http.ResponseWriter, req *http.Request) {
	form := &dto.PostCreate{}

	err := json.NewDecoder(req.Body).Decode(form)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	user_id := req.Context().Value(middleware.ContextUserIDKey)
	if user_id == nil {
		utils.ResponseError(res, errors.New("no jwt provided"), http.StatusUnauthorized)
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

func ToggleLikePost(res http.ResponseWriter, req *http.Request) {
	form := &dto.LikePost{}

	err := json.NewDecoder(req.Body).Decode(form)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	user_id := req.Context().Value(middleware.ContextUserIDKey)
	if user_id == nil {
		utils.ResponseError(res, errors.New("no jwt provided"), http.StatusUnauthorized)
		return
	}

	model, err := models.GetLike(form.PostID, user_id.(int))
	if err == nil {
		models.DeleteLike(model.ID)
		utils.ResponseEmptySucess(res)
		return
	}

	model = mappers.LikePostToLike(form, user_id.(int))

	_, err = model.Save()
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	like, err := mappers.ToLikePost(model)
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(like)
	res.Write(b)
}
