package models

import (
	"api/db"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int       `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	IsOnline   bool      `json:"isOnline,omitempty"`
	Status     string    `json:"status,omitempty"`
	BirthDate  time.Time `json:"birthDate,omitempty"`
	City       string    `json:"city,omitempty"`
	Avatar_ID  int       `json:"avatar_id,omitempty"`
	created_at time.Time `json:"user_created_at,omitempty`
	deleted_at time.Time `json:"user_created_at,omitempty`
}

// Validate function for validate
func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") && !strings.Contains(user.Email, ".") {
		return map[string]interface{}{"status": "invalid", "message": "Email address format is incorrect"}, false
	}

	if len(user.Password) < 6 {
		return map[string]interface{}{"status": "invalid", "message": "Password is minimum 6 character"}, false
	}
	temp := &User{}
	sql := fmt.Sprintf("SELECT id,email,username FROM users	WHERE email = '%s'", user.Email)
	data, err := db.DB.Query(sql)
	if err != nil {
		return map[string]interface{}{"status": "invalid", "message": "Something went wrong, please contact admin or developer."}, false
	}
	for data.Next() {
		err = data.Scan(&temp.ID, &temp.Email, &temp.Username)
		if err != nil {
			saveError := fmt.Sprintf("Error Looping data, and %s", err)
			return map[string]interface{}{"status": "invalid", "message": saveError}, false
		}
	}
	if temp.Email != "" {
		return map[string]interface{}{"status": "invalid", "message": "Email address already in use by another user."}, false
	}

	return map[string]interface{}{"status": "Valid", "message": "Requirement passed"}, true
}

func (user *User) Register() map[string]interface{} {
	if rsp, status := user.Validate(); !status {
		return rsp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	now := time.Now().Format("2006-01-02 15:04:05")
	idUser := 0
	sql := fmt.Sprintf(`INSERT INTO users (email, username, password, created_at, updated_at, status) 
						VALUES ('%s', '%s', '%s', '%s', '%s', '1')
						RETURNING id ; `, user.Email, user.Username, user.Password, now, now)
	err := db.DB.QueryRow(sql).Scan(&idUser)
	if err != nil || idUser == 0 {
		return map[string]interface{}{"status": "invalid", "message": "Insert Errors call admin or developer "}
	}
	Hours := 24
	Mins := 12
	timein := time.Now().Local().Add(time.Hour*time.Duration(Hours) +
		time.Minute*time.Duration(Mins))

	tk := &Token{UserID: uint(idUser), Email: user.Email, TimeExp: timein}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("asdf"))

	return map[string]interface{}{"status": "valid", "message": "Account is successfully created ", "token": tokenString}
}

// Login function for checking login valid or not
func (user *User) Login() map[string]interface{} {
	sql := fmt.Sprintf("SELECT id,email,username,password FROM users WHERE email = '%s'", user.Email)
	data, err := db.DB.Query(sql)
	if err != nil {
		log.Println("error query : ", err)
	}
	temp := &User{}
	for data.Next() {
		err = data.Scan(&temp.ID, &temp.Email, &temp.Username, &temp.Password)
		if err != nil {
			saveError := fmt.Sprintf("Error Looping data, and %s", err)
			log.Println(saveError)
		}
	}
	if temp.Email == "" {
		return map[string]interface{}{"status": "invalid", "message": "Email Invalid please try again."}
	}

	err = bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(user.Password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"status": "invalid", "message": "Password Invalid."}
	}
	Hours := 24
	Mins := 12
	timein := time.Now().Local().Add(time.Hour*time.Duration(Hours) +
		time.Minute*time.Duration(Mins))

	tk := &Token{UserID: uint(temp.ID), Email: temp.Email, TimeExp: timein}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("asdf"))

	return map[string]interface{}{"status": "valid", "message": "Login is Success", "token": tokenString}
}
