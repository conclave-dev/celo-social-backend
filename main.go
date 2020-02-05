package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/stella-zone/celo-social-backend/routes"
)

func main() {
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)

	log.Printf("Starting at 8081")
	router := setUpRouter()
	log.Fatal(http.ListenAndServe("localhost:8081", cors(router)))
}

func setUpRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes.SetUpRoutes(router)
	return router
}
