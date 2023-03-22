package services

import (
	"encoding/json"
	"net/http"
)

func newResponse(messageType string, message string, data interface{}) response {
	return response{
		messageType,
		message,
		data,
	}
}

func responseJSON(w http.ResponseWriter, statusCode int, response response) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
