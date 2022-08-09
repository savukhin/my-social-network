package app

import (
	"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"api/app/auth"
	"log"
	"net/http"
)

func Launch() {
	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()
	router = auth.Routes(router.PathPrefix("/auth").Subrouter())

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(":4201", handlers.CORS(credentials, methods, origins)(router)))
	fmt.Println("Server started")
}
