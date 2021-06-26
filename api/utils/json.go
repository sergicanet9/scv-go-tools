package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON makes the response with payload as json format
func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// ResponseError makes the error response with payload as json format
func ResponseError(w http.ResponseWriter, status int, message string) {
	ResponseJSON(w, status, map[string]string{"error": message})
}
