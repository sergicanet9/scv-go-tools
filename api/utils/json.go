package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ResponseJSON makes the response with payload as json format
func ResponseJSON(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		log.Print(fmt.Sprintf("REQUEST %s:%s, RESPONSE %d: %s", r.Method, r.URL.String(), http.StatusInternalServerError, err.Error()))
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			w.Header().Set(name, value)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))

	log.Print(fmt.Sprintf("REQUEST %s:%s, RESPONSE %d: %s", r.Method, r.URL.String(), status, string(response)))
}

// ResponseError makes the error response with payload as json format
func ResponseError(w http.ResponseWriter, r *http.Request, status int, message string) {
	ResponseJSON(w, r, status, map[string]string{"error": message})
}
