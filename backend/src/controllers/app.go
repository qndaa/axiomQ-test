package controllers

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/qndaa/backend/src/middlewares"
	"log"
	"net/http"
)

type App struct {
	r *mux.Router
}

func (a *App) Init() {
	a.r = mux.NewRouter()
	a.r.HandleFunc("/api/book", BookControllerInstance.FetchAll).Methods("GET")
	a.r.HandleFunc("/api/book", BookControllerInstance.Create).Methods("POST")
	a.r.HandleFunc("/api/book/{id}", BookControllerInstance.FetchOne).Methods("GET")
	a.r.HandleFunc("/api/login", AuthInstanceController.Login).Methods("POST")
}

func (a *App) StartServer() {
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), handlers.CORS(headers, methods, origins)(a.GetRouter())))
}

func (a *App) GetRouter() *mux.Router {
	return a.r
}

func (a *App) ApplyMiddlewares() {
	a.GetRouter().Use(middlewares.IsAuthenticated)
}
