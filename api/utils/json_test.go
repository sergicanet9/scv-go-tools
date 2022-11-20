package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestResponseJSON_OK checks that ResponseJSON returns the expected response when all goes as expected
func TestResponseJSON_OK(t *testing.T) {
	// Arrange
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("test-header", "test")
	body := map[string]string{"body": "test-body"}
	expectedResponse := map[string]string{"response": "test-response"}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseJSON(w, r, []byte(fmt.Sprintf("%v", body)), http.StatusOK, expectedResponse)
	})

	// Act
	handlerToTest.ServeHTTP(rr, req)

	// Assert
	if want, got := http.StatusOK, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

// TestResponseJSON_PayloadNotMarshalled checks that ResponseJSON returns the expected response when the response cannot be marshalled
func TestResponseJSON_PayloadNotMarshalled(t *testing.T) {
	// Arrange
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	notMarshableResponse := map[string]interface{}{"response": make(chan int)}
	expectedResponse := map[string]string(map[string]string{"error": "json: unsupported type: chan int"})

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseJSON(w, r, nil, http.StatusOK, notMarshableResponse)
	})

	// Act
	handlerToTest.ServeHTTP(rr, req)

	// Assert
	if want, got := http.StatusOK, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

// TestResponseError_Ok checks that ResponseError returns the expected error response when all goes as expected
func TestResponseError_Ok(t *testing.T) {
	// Arrange
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	err := "test-error"
	expectedResponse := map[string]string{"error": err}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, r, nil, http.StatusBadRequest, err)
	})

	// Act
	handlerToTest.ServeHTTP(rr, req)

	// Assert
	if want, got := http.StatusBadRequest, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}
