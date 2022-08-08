package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

// "fmt"
// "log"
// "net/http"
// "time"

// "github.com/cakazies/go-postgre/application/api"
// "github.com/cakazies/go-postgre/application/middleware"
// "github.com/gorilla/mux"
// "github.com/spf13/viper"

func Routes(router *mux.Router) {
	// engine.GET("/users/:id", controllable.GetUser)
	router.HandleFunc("login/", Login).Methods(http.MethodPost)
	router.HandleFunc("register/", Register).Methods(http.MethodPost)
}
