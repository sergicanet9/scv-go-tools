package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseJSON(t *testing.T) {
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	body := map[string]string{"body": "test-body"}
	expectedResponse := map[string]string{"response": "test-response"}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseJSON(w, r, []byte(fmt.Sprintf("%v", body)), http.StatusOK, expectedResponse)
	})
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusOK, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

func TestResponseError(t *testing.T) {
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	body := map[string]string{"body": "test-body"}
	err := "test-error"
	expectedResponse := map[string]string{"error": err}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, r, []byte(fmt.Sprintf("%v", body)), http.StatusOK, err)
	})
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusOK, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}
