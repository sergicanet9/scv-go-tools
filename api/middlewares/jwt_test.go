package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var (
	headerName     = "Authorization"
	secret         = "test-secret"
	jwtOk          = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.mpHl842O7xEZjgQ8CyX8xYLDoEORGVMnAxULkW-u8Ek"
	urlJWT         = "http://testing"
	handlerFuncJWT = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
)

func TestJWTOk(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, urlJWT, nil)
	req.Header.Add(headerName, jwtOk)

	handlerToTest := JWT(handlerFuncJWT, secret, jwt.MapClaims{})
	handlerToTest.ServeHTTP(rr, req)

	if want, got := http.StatusOK, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}
}

func TestJWTMissingHeader(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, urlJWT, nil)
	expectedResponse := map[string]string{"error": "an authorization header is required"}

	handlerToTest := JWT(handlerFuncJWT, secret, jwt.MapClaims{})
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

func TestJWTMalformedToken(t *testing.T) {
	jwtMalformed := "123"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, urlJWT, nil)
	req.Header.Add(headerName, jwtMalformed)
	expectedResponse := map[string]string{"error": "authorization header not properly formated, should be Bearer + {token}"}

	handlerToTest := JWT(handlerFuncJWT, secret, jwt.MapClaims{})
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

func TestJWTInvalidSecret(t *testing.T) {
	jwtInvalidSecret := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.df4nTfuWWdndfrlIxF0iWUrrcANrM4bzKdbYa9VeAj8"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, urlJWT, nil)
	req.Header.Add(headerName, jwtInvalidSecret)
	expectedResponse := map[string]string{"error": "signature is invalid"}

	handlerToTest := JWT(handlerFuncJWT, secret, jwt.MapClaims{})
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

func TestJWTMissingClaim(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, urlJWT, nil)
	req.Header.Add(headerName, jwtOk)
	expectedResponse := map[string]string{"error": "required claim test-claim not found or incorrect"}

	handlerToTest := JWT(handlerFuncJWT, secret, jwt.MapClaims{"test-claim": true})
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
