package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spintoaguero/meli-challenge/configs"
	"github.com/spintoaguero/meli-challenge/internal/services"
)

func SetupRoutes(database *configs.Database) {

	muxRouter := mux.NewRouter()

	managerHandler := &services.ManagerHandler{
		Database: database,
	}

	muxRouter.HandleFunc("/hello", services.Hello).Methods(http.MethodGet)

	muxRouter.HandleFunc("/headers", services.Headers).Methods(http.MethodGet)

	muxRouter.HandleFunc("/generate", managerHandler.GenerateShortUrl).Methods(http.MethodGet)

	routesHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}).Handler(muxRouter)

	http.ListenAndServe(":8080", routesHandler)
}
