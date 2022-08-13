package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID  int       `json:"user_id,omitemp"`
	Email   string    `json:"email,omitemp"`
	TimeExp time.Time `json:"time_exp,omitemp"`
	jwt.StandardClaims
}
