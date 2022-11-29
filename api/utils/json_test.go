package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/sergicanet9/scv-go-tools/v3/wrappers"
	"github.com/stretchr/testify/assert"
)

type targetType struct {
	TestDuration1 Duration
	TestDuration2 Duration
}

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
func TestResponseError_BadRequestNonExistent(t *testing.T) {
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
	if want, got := http.StatusBadRequest, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

// TestResponseError_BadRequest checks that ResponseError returns the expected error response when the received error is of type ValidationErr
func TestResponseError_BadRequestNonValid(t *testing.T) {
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

// TestLoadJSON_Ok checks that LoadJSON does not return an error an parses Duration object properly when all goes as expected
func TestLoadJSON_Ok(t *testing.T) {
	// Arrange
	target := targetType{}

	_, filePath, _, _ := runtime.Caller(0)
	dir, err := os.MkdirTemp(filepath.Join(filePath, "../../.."), "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file, err := os.CreateTemp(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	bytes := []byte(`{"TestDuration1":10,"TestDuration2":"10s"}`)
	err = os.WriteFile(file.Name(), bytes, 0644)
	if err != nil {
		t.Fatal(err)
	}

	expectedDuration1 := Duration{time.Duration(10)}
	duration2, _ := time.ParseDuration("10s")
	expectedDuration2 := Duration{duration2}

	// Act
	err = LoadJSON(file.Name(), &target)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedDuration1, target.TestDuration1)
	assert.Equal(t, expectedDuration2, target.TestDuration2)
}

// TestLoadJSON_NonExistentFile checks that LoadJSON returns an error when the specified file does not exist
func TestLoadJSON_NonExistentFile(t *testing.T) {
	// Arrange
	target := targetType{}
	expectedError := "ignoring config file : stat : no such file or directory"

	// Act
	err := LoadJSON("", &target)

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// TestLoadJSON_FileNotAccessible checks that LoadJSON returns an error when the specified file is not accessible
func TestLoadJSON_FileNotAccessible(t *testing.T) {
	// Arrange
	target := targetType{}

	_, filePath, _, _ := runtime.Caller(0)
	dir, err := os.MkdirTemp(filepath.Join(filePath, "../../.."), "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file, err := os.CreateTemp(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	err = file.Chmod(os.ModeExclusive)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := fmt.Sprintf("error opening file %s: open %s: permission denied", file.Name(), file.Name())

	// Act
	err = LoadJSON(file.Name(), &target)

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// TestLoadJSON_FileIsADirectory checks that LoadJSON returns an error when the specified file is a directory
func TestLoadJSON_FileIsADirectory(t *testing.T) {
	// Arrange
	target := targetType{}

	_, filePath, _, _ := runtime.Caller(0)
	dir, err := os.MkdirTemp(filepath.Join(filePath, "../../.."), "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	expectedError := fmt.Sprintf("error reading file %s: read %s: is a directory", dir, dir)

	// Act
	err = LoadJSON(dir, &target)

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// TestLoadJSON_InvalidDurationStringFormat checks that LoadJSON returns an error when the file contains a string that cannot be parsed to a Duration object
func TestLoadJSON_InvalidDurationStringFormat(t *testing.T) {
	// Arrange
	target := targetType{}

	_, filePath, _, _ := runtime.Caller(0)
	dir, err := os.MkdirTemp(filepath.Join(filePath, "../../.."), "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file, err := os.CreateTemp(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	bytes := []byte(`{"TestDuration1":"10secs"}`)
	err = os.WriteFile(file.Name(), bytes, 0644)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := fmt.Sprintf("error unmarshaling file %s: time: unknown unit \"secs\" in duration \"10secs\"", file.Name())

	// Act
	err = LoadJSON(file.Name(), &target)

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// TestLoadJSON_InvalidDurationType checks that LoadJSON returns an error when the file contains a type that cannot be parsed to a Duration object
func TestLoadJSON_InvalidDurationType(t *testing.T) {
	// Arrange
	target := targetType{}

	_, filePath, _, _ := runtime.Caller(0)
	dir, err := os.MkdirTemp(filepath.Join(filePath, "../../.."), "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file, err := os.CreateTemp(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	bytes := []byte(`{"TestDuration1":{"Test":1}}`)
	err = os.WriteFile(file.Name(), bytes, 0644)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := fmt.Sprintf("error unmarshaling file %s: invalid duration", file.Name())

	// Act
	err = LoadJSON(file.Name(), &target)

	// Assert
	assert.Equal(t, expectedError, err.Error())
}
