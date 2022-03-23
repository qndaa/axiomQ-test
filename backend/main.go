package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/qndaa/backend/src/controllers"
	"github.com/qndaa/backend/src/db"
	"github.com/qndaa/backend/src/middlewares"
	"log"
	"net/http"
)

func run() {
	db.Connect()
	db.AutoMigration()

	r := mux.NewRouter()
	r.HandleFunc("/api/book", controllers.BookControllerInstance.FetchAll).Methods("GET")
	r.HandleFunc("/api/book", controllers.BookControllerInstance.Create).Methods("POST")
	r.HandleFunc("/api/book/{id}", controllers.BookControllerInstance.FetchOne).Methods("GET")
	r.HandleFunc("/api/login", controllers.AuthInstanceController.Login).Methods("POST")
	r.Use(middlewares.IsAuthenticated)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), handlers.CORS(headers, methods, origins)(r)))
}

func main() {
	run()
}
