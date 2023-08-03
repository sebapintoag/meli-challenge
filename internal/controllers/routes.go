package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spintoaguero/meli-challenge/internal/controllers/handlers"
	"github.com/spintoaguero/meli-challenge/internal/repositories"
	"github.com/spintoaguero/meli-challenge/pkg/utils"
)

func SetupRoutes(linkRepository *repositories.LinkRepository) {

	// Create new link handler
	linkHandler := handlers.NewLinkHandler(linkRepository)

	// Create router
	muxRouter := mux.NewRouter()
	muxRouter.StrictSlash(true)

	// Set '/api/v1' subpath for API REST endpoints
	apiV1 := muxRouter.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/create", linkHandler.CreateShortUrl).Methods(http.MethodPost)
	apiV1.HandleFunc("/find", linkHandler.FindUrl).Methods(http.MethodPost)
	apiV1.HandleFunc("/delete", linkHandler.DeleteShortUrl).Methods(http.MethodDelete)

	// Set handler for short URL redirection
	muxRouter.HandleFunc("/{key}", linkHandler.Redirect).Methods(http.MethodGet)

	muxRouter.Use(utils.ContentTypeApplicationJsonMiddleware)

	routesHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:8080", "http://localhost:3000", "http://localhost:40", "http://me.li", "http://me.li:40"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	}).Handler(muxRouter)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("MELI_APP_PORT")), routesHandler)
}
