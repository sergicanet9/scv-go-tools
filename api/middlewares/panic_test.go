package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicRecover_NoPanic(t *testing.T) {
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := PanicRecover(handlerFunc)
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusOK, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}
}

func TestPanicRecover_Panic(t *testing.T) {
	var url = "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	expectedResponse := map[string]string{"error": "test panic"}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})
	handlerToTest := PanicRecover(handlerFunc)
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusInternalServerError, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}
