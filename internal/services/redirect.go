package services

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Hello(w http.ResponseWriter, req *http.Request) {

	fmt.Println("asdasdsadsadsadsa")
	id := uuid.NewString()
	fmt.Fprintf(w, id)
}
