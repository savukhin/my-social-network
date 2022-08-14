package models

import (
	"api/db"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int            `json:"id,omitempty"`
	Username  string         `json:"username,omitempty"`
	Name      string         `json:"name,omitempty"`
	Email     string         `json:"email,omitempty"`
	Password  string         `json:"password,omitempty"`
	Password2 string         `json:"password2,omitempty"`
	IsOnline  bool           `json:"isOnline,omitempty"`
	Status    sql.NullString `json:"status,omitempty"`
	BirthDate sql.NullTime   `json:"birthDate,omitempty"`
	City      sql.NullString `json:"city,omitempty"`
	Avatar_ID sql.NullInt64  `json:"avatar_id,omitempty"`
	CreatedAt time.Time      `json:"user_created_at,omitempty"`
	DeletedAt time.Time      `json:"user_deleted_at,omitempty"`
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") && !strings.Contains(user.Email, ".") {
		return map[string]interface{}{"status": "error", "message": "Email address format is incorrect"}, false
	}

	if len(user.Password) < 6 {
		return map[string]interface{}{"status": "error", "message": "Password is minimum 6 character"}, false
	}
	temp := &User{}
	sql := fmt.Sprintf("SELECT id,email,username FROM users	WHERE email = '%s'", user.Email)
	data, err := db.DB.Query(sql)
	if err != nil {
		return map[string]interface{}{"status": "error", "message": "Something went wrong, please contact admin or developer."}, false
	}
	for data.Next() {
		err = data.Scan(&temp.ID, &temp.Email, &temp.Username)
		if err != nil {
			saveError := fmt.Sprintf("Error Looping data, and %s", err)
			return map[string]interface{}{"status": "error", "message": saveError}, false
		}
	}
	if temp.Email != "" {
		return map[string]interface{}{"status": "error", "message": "Email address already in use by another user."}, false
	}

	return map[string]interface{}{"status": "ok", "message": "Requirement passed"}, true
}

func (user *User) Register() map[string]interface{} {
	if rsp, status := user.Validate(); !status {
		return rsp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	idUser := 0
	sql := fmt.Sprintf(`INSERT INTO users (email, username, name, password) 
						VALUES ('%s', '%s', '%s', '%s')
						RETURNING id ; `, user.Email, user.Username, user.Username, user.Password)
	err := db.DB.QueryRow(sql).Scan(&idUser)
	if err != nil || idUser == 0 {
		return map[string]interface{}{"status": "error", "message": "Insert Errors call admin or developer "}
	}
	Hours := 24
	Mins := 12
	timein := time.Now().Local().Add(time.Hour*time.Duration(Hours) +
		time.Minute*time.Duration(Mins))

	tk := &Token{UserID: idUser, Email: user.Email, TimeExp: timein}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("asdf"))

	return map[string]interface{}{"status": "valid", "message": "Account is successfully created ", "id_token": tokenString, "expires_at": timein}
}

func (user *User) Login() map[string]interface{} {
	sql := fmt.Sprintf("SELECT id, email, password FROM users WHERE username = '%s'", user.Username)
	row := db.DB.QueryRow(sql)

	temp := &User{}
	err := row.Scan(&temp.ID, &temp.Email, &temp.Password)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "No such user with this username",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(user.Password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"status": "error", "message": "Password Invalid."}
	}

	Hours := 24
	Mins := 12
	timein := time.Now().Local().Add(time.Hour*time.Duration(Hours) +
		time.Minute*time.Duration(Mins))

	tk := &Token{UserID: temp.ID, Email: temp.Email, TimeExp: timein}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("asdf"))

	return map[string]interface{}{"status": "ok", "message": "Login is Success", "id_token": tokenString, "expires_at": timein}
}

func (user *User) GetProfile() (*User, error) {
	sql := fmt.Sprintf("SELECT id, username, name, email, status, city, birthdate, avatar_id, isOnline FROM users WHERE id = %d", user.ID)
	row := db.DB.QueryRow(sql)

	temp := &User{}
	err := row.Scan(
		&temp.ID, &temp.Username, &temp.Name, &temp.Email,
		&temp.Status, &temp.City, &temp.BirthDate, &temp.Avatar_ID, &temp.IsOnline,
	)

	if err != nil {
		return nil, err
	}

	return temp, nil
}

func (user *User) ChangeProfile() (interface{}, error) {
	temp, err := GetUserByID(user.ID)

	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "No such user with this id",
		}, err
	}

	return temp, nil
}
