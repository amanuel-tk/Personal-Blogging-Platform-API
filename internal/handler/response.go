package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode response:%v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJson(w, status, errorResponse{
		Error: message,
	})
}
