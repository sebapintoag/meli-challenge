package utils

import (
	"net/http"
)

// Middleware type definition
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Applies middlewares to a http.HandlerFunc
func Chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
