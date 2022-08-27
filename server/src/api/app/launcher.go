package app

import (
	"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"api/app/auth"
	"api/app/chat"
	"api/app/posts"
	"api/app/users"
	"log"
	"net/http"
)

func Launch() {
	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()
	auth.Routes(router.PathPrefix("/auth").Subrouter())
	users.Routes(router.PathPrefix("/users").Subrouter())
	chat.Routes(router.PathPrefix("").Subrouter())
	posts.Routes(router.PathPrefix("/posts/").Subrouter())

	hub := chat.CreateHub()
	go hub.Run()

	r.HandleFunc("/ws/chat_id={chat_id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	r.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("works"))
	})

	// router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "enctype"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(
		":4201",
		handlers.CORS(credentials, methods, origins, ttl, headers)(r),
		// router,
	))
	fmt.Println("Server started")
}
