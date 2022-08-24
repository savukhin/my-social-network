package posts

import (
	"api/db/models"
	"api/dto"
	"api/mappers"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserPosts(res http.ResponseWriter, req *http.Request) {
	user_id, err := strconv.Atoi(mux.Vars(req)["user_id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	contents, err := models.GetPostsByUserID(user_id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
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
