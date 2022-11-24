package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sergicanet9/scv-go-tools/v3/wrappers"
	"github.com/stretchr/testify/assert"
)

// TestResponseJSON_Ok checks that ResponseJSON returns the expected response when all goes as expected
func TestResponseJSON_Ok(t *testing.T) {
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

// TestResponseError_NoContent checks that ResponseError returns the expected error response when the received error is of type NonExistentErr
func TestResponseError_NoContent(t *testing.T) {
	// Arrange
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	err := wrappers.NonExistentErr
	expectedResponse := map[string]string{"error": err.Error()}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, r, nil, err)
	})

	// Act
	handlerToTest.ServeHTTP(rr, req)

	// Assert
	if want, got := http.StatusNoContent, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

// TestResponseError_BadRequest checks that ResponseError returns the expected error response when the received error is of type ValidationErr
func TestResponseError_BadRequest(t *testing.T) {
	// Arrange
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	err := wrappers.ValidationErr
	expectedResponse := map[string]string{"error": err.Error()}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, r, nil, err)
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

// TestResponseError_Unauthorized checks that ResponseError returns the expected error response when the received error is of type UnauthorizedErr
func TestResponseError_Unauthorized(t *testing.T) {
	// Arrange
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	err := wrappers.UnauthorizedErr
	expectedResponse := map[string]string{"error": err.Error()}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, r, nil, err)
	})

	// Act
	handlerToTest.ServeHTTP(rr, req)

	// Assert
	if want, got := http.StatusUnauthorized, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

// TestResponseError_RequestTimeout checks that ResponseError returns the expected error response when the received error is of type DeadlineExceeded
func TestResponseError_RequestTimeout(t *testing.T) {
	// Arrange
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	err := context.DeadlineExceeded
	expectedResponse := map[string]string{"error": err.Error()}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, r, nil, err)
	})

	// Act
	handlerToTest.ServeHTTP(rr, req)

	// Assert
	if want, got := http.StatusRequestTimeout, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

// TestResponseError_InternalServerError checks that ResponseError returns the expected error response when the received error is not any of the previous types
func TestResponseError_InternalServerError(t *testing.T) {
	// Arrange
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	err := errors.New("test")
	expectedResponse := map[string]string{"error": err.Error()}

	handlerToTest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, r, nil, err)
	})

	// Act
	handlerToTest.ServeHTTP(rr, req)

	// Assert
	if want, got := http.StatusInternalServerError, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}
