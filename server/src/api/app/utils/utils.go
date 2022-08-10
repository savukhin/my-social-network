package utils

import (
	"api/db/models"
	"bytes"
	"encoding/json"
	"errors"
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
