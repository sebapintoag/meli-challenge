package configs

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spintoaguero/meli-challenge/internal/services/shortener"
)

func SetupRoutes() {

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/hello", shortener.Hello).Methods(http.MethodGet)

	muxRouter.HandleFunc("/headers", shortener.Headers).Methods(http.MethodGet)

	muxRouter.HandleFunc("/generate", shortener.GenerateShortUrl).Methods(http.MethodGet)

	routesHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}).Handler(muxRouter)

	http.ListenAndServe(":8080", routesHandler)
}
