package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID  int       `json:"user_id,omitempty"`
	Email   string    `json:"email,omitempty"`
	TimeExp time.Time `json:"time_exp,omitempty"`
	jwt.StandardClaims
}
