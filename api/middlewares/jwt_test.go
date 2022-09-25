package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestJWT_Ok(t *testing.T) {
	jwtOk := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.mpHl842O7xEZjgQ8CyX8xYLDoEORGVMnAxULkW-u8Ek"
	secret := "test-secret"
	headerName := "Authorization"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(headerName, jwtOk)

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{})
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusOK, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}
}

func TestJWT_MissingHeader(t *testing.T) {
	secret := "test-secret"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	expectedResponse := map[string]string{"error": "an authorization header is required"}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{})
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusUnauthorized, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

func TestJWT_MalformedToken(t *testing.T) {
	jwtMalformed := "123"
	secret := "test-secret"
	headerName := "Authorization"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(headerName, jwtMalformed)
	expectedResponse := map[string]string{"error": "authorization header not properly formated, should be Bearer + {token}"}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{})
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusUnauthorized, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

func TestJWT_InvalidSecret(t *testing.T) {
	jwtInvalidSecret := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.df4nTfuWWdndfrlIxF0iWUrrcANrM4bzKdbYa9VeAj8"
	secret := "test-secret"
	headerName := "Authorization"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(headerName, jwtInvalidSecret)
	expectedResponse := map[string]string{"error": "signature is invalid"}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{})
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusUnauthorized, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)
}

func TestJWT_MissingClaim(t *testing.T) {
	jwtMissingClaim := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.mpHl842O7xEZjgQ8CyX8xYLDoEORGVMnAxULkW-u8Ek"
	secret := "test-secret"
	headerName := "Authorization"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(headerName, jwtMissingClaim)
	expectedResponse := map[string]string{"error": "required claim test-claim not found or incorrect"}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{"test-claim": true})
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusUnauthorized, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error parsing the response: %s", err)
	}
	assert.Equal(t, expectedResponse, response)

}
