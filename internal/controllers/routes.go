package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spintoaguero/meli-challenge/internal/controllers/admin"
	"github.com/spintoaguero/meli-challenge/internal/controllers/redirect"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
)

func SetupRoutes(database *mongodb.Database) {

	muxRouter := mux.NewRouter()

	adminHandler := &admin.AdminHandler{
		Database: database,
	}

	muxRouter.HandleFunc("/hello", redirect.Hello).Methods(http.MethodGet)

	muxRouter.HandleFunc("/headers", admin.Headers).Methods(http.MethodGet)

	muxRouter.HandleFunc("/generate", adminHandler.GenerateShortUrl).Methods(http.MethodGet)

	routesHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}).Handler(muxRouter)

	http.ListenAndServe(":8080", routesHandler)
}
