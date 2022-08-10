package app

import (
	"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"api/app/auth"
	"api/app/users"
	"log"
	"net/http"
)

func Launch() {
	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()
	auth.Routes(router.PathPrefix("/auth").Subrouter())
	users.Routes(router.PathPrefix("/users").Subrouter())

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(
		":4201",
		handlers.CORS(credentials, methods, origins, ttl, headers)(router),
		// router,
	))
	fmt.Println("Server started")
}
