package handler

import (
	"encoding/json"
	"net/http"
)

func writeErrorJSON(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": "request failed",
	})
}