package middleware

import (
	"api/db/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ContextKey string

const ContextUserIDKey ContextKey = "user_id"

// JwtAuthentication for JWT
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenHeader := req.Header.Get("Authorization")

		if tokenHeader == "" {
			rsp := map[string]interface{}{"status": "invalid", "message": "Token is not Present ;"}
			res.Header().Add("Content-Type", "application/json")
			json.NewEncoder(res).Encode(rsp)
			return
		}

		tk := &models.Token{}
		tokenValue := tokenHeader
		token, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("asdf"), nil
		})

		if err != nil {
			rsp := map[string]interface{}{"status": "invalid", "message": "Malformed Authentication Token Please Login Again;"}
			res.Header().Add("Content-Type", "application/json")
			json.NewEncoder(res).Encode(rsp)
			return
		}

		// check for time expired
		diff := tk.TimeExp.Sub(time.Now())
		if diff < 0 {
			rsp := map[string]interface{}{"status": "invalid", "message": "Time Expired, please login again;"}
			res.Header().Add("Content-Type", "application/json")
			json.NewEncoder(res).Encode(rsp)
			return
		}

		if !token.Valid {
			rsp := map[string]interface{}{"status": "invalid", "message": "Invalid/Format Auth Token ;"}
			res.Header().Add("Content-Type", "application/json")
			json.NewEncoder(res).Encode(rsp)
			return
		}

		ctx := context.WithValue(req.Context(), ContextUserIDKey, tk.UserID)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}
