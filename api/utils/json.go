package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseJSON makes the response with payload as json format
func ResponseJSON(w http.ResponseWriter, r *http.Request, body []byte, status int, payload interface{}) {
	for name, values := range r.Header {
		if name != "Content-Length" {
			for _, value := range values {
				w.Header().Set(name, value)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Response paylod could not be marshalled")
		response, _ = json.Marshal(map[string]string{"error": err.Error()})
	}

	w.Write([]byte(response))
	log.Printf("REQUEST %s:%s", r.Method, r.URL.String())
	if body != nil {
		log.Printf("BODY: %s", string(body))
	}
	log.Printf("RESPONSE %d: %s", status, string(response))
}

// ResponseError makes the error response with payload as json format
func ResponseError(w http.ResponseWriter, r *http.Request, body []byte, status int, message string) {
	ResponseJSON(w, r, body, status, map[string]string{"error": message})
}
