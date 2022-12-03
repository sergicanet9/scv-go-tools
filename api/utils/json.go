package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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
	if errors.Is(err, wrappers.NonExistentErr) || errors.Is(err, wrappers.ValidationErr) {
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

// LoadJSON opens the specified file and unmarshals its JSON content in the received struct
func LoadJSON(filePath string, target interface{}) error {
	if _, err := os.Stat(filePath); err != nil {
		return fmt.Errorf("ignoring config file %v: %w", filePath, err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %v: %w", filePath, err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file %v: %w", filePath, err)
	}

	err = json.Unmarshal(byteValue, target)
	if err != nil {
		return fmt.Errorf("error unmarshaling file %v: %w", filePath, err)
	}

	return nil
}

// Duration allows to unmarshal time into time.Duration
// https://github.com/golang/go/issues/10275
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	var v interface{}
	json.Unmarshal(b, &v)

	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
	case string:
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("invalid duration")
	}
	return nil
}
