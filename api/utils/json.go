package utils

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/sergicanet9/scv-go-tools/v3/wrappers"
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
// Calls ResponseJSON with different error codes depending of the type of the error received
func ResponseError(w http.ResponseWriter, r *http.Request, body []byte, err error) {
	var status int
	if errors.Is(err, wrappers.NonExistentErr) {
		status = http.StatusNoContent
	} else if errors.Is(err, wrappers.ValidationErr) {
		status = http.StatusBadRequest
	} else if errors.Is(err, wrappers.UnauthorizedErr) {
		status = http.StatusUnauthorized
	} else if errors.Is(err, context.DeadlineExceeded) {
		status = http.StatusRequestTimeout
	} else {
		status = http.StatusInternalServerError
	}
	ResponseJSON(w, r, body, status, map[string]string{"error": err.Error()})
}
