package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessBody struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data,omitempty"`
}

type ErrorBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ErrorResponse(w http.ResponseWriter, req *http.Request, status string, code int, err error) {
	response := ErrorBody{
		Status:  status,
		Message: err.Error(),
	}
	marshaledResponse, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledResponse)
}

func SuccessResponse(w http.ResponseWriter, req *http.Request, body map[string]interface{}, code int) {
	response := SuccessBody{
		Status: "success",
		Data:   body,
	}
	marshaledResponse, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledResponse)
}
