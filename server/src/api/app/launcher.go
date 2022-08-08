package app

import (
	"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	// "github.com/rs/cors"

	"api/app/auth"
	"api/app/utils"
	"log"
	"net/http"
)

func Launch() {
	r := mux.NewRouter()
	router := r.PathPrefix("api/").Subrouter()
	// router.HandleFunc("/hello-world", helloWorld)
	auth.Routes(router.PathPrefix("users").Subrouter())

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET"})
	// ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	// Solves Cross Origin Access Issue
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"http://localhost:4200"},
	// })
	// handler := c.Handler(r)

	// srv := &http.Server{
	// 	Handler: handler,
	// 	Addr:    ":4201",
	// }

	// log.Fatal(srv.ListenAndServe())
	log.Fatal(http.ListenAndServe(":4201", handlers.CORS(credentials, methods, origins)(router)))
	fmt.Println("Starting...")
	// log.Fatal(http.ListenAndServe(":4201", router))
	fmt.Println("Server started")
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new request")
	// setupCorsResponse(&w, r)
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}
