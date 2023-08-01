package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spintoaguero/meli-challenge/internal/controllers/api/v1/admin"
	"github.com/spintoaguero/meli-challenge/internal/controllers/redirect"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"github.com/spintoaguero/meli-challenge/pkg/utils"
)

func SetupRoutes(database *mongodb.Database) {

	muxRouter := mux.NewRouter()
	muxRouter.StrictSlash(true)

	adminHandler := &admin.AdminHandler{
		Database: database,
	}

	redirectHandler := &redirect.RedirectHandler{
		Database: database,
	}

	apiV1 := muxRouter.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/create", adminHandler.CreateShortUrl).Methods(http.MethodPost)
	apiV1.HandleFunc("/find", adminHandler.FindUrl).Methods(http.MethodPost)
	apiV1.HandleFunc("/delete", adminHandler.DeleteShortUrl).Methods(http.MethodDelete)

	muxRouter.HandleFunc("/{key}", redirectHandler.Perform).Methods(http.MethodGet)

	muxRouter.Use(utils.ContentTypeApplicationJsonMiddleware)

	routesHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:8080", "http://localhost:3000", "http://localhost:40"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}).Handler(muxRouter)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("MELI_APP_PORT")), routesHandler)
}
