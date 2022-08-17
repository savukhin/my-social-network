package utils

import (
	"api/db/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func StructToJSON(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func UnpackJWT(tokenHeader string) (*models.Token, error) {
	if tokenHeader == "" {
		return nil, errors.New("token is not present")
	}

	tk := &models.Token{}
	tokenValue := tokenHeader
	token, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte("asdf"), nil
	})

	if err != nil {
		return nil, errors.New("malformed authentication token")
	}

	diff := time.Until(tk.TimeExp)

	if diff < 0 {
		return nil, errors.New("token expired")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return tk, nil
}

func ResponseEmptySucess(res http.ResponseWriter) {
	b, _ := json.Marshal(map[string]interface{}{"success": "success"})

	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func ResponseError(res http.ResponseWriter, message error, status int) error {
	b, err := json.Marshal(map[string]interface{}{"error": message.Error()})
	if err != nil {
		return err
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(b)
	return nil
}
