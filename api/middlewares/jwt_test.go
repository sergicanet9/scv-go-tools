package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

// TestJWT_Ok checks that the middleware does not return an error when the token matches with all requirements
func TestJWT_Ok(t *testing.T) {
	// Arrange
	jwtOk := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.mpHl842O7xEZjgQ8CyX8xYLDoEORGVMnAxULkW-u8Ek"
	secret := "test-secret"
	headerName := "Authorization"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(headerName, jwtOk)

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{})

	// Act
	handlerToTest.ServeHTTP(rr, req)

	// Assert
	if want, got := http.StatusOK, rr.Code; want != got {
		t.Fatalf("unexpected http status code: want=%d but got=%d", want, got)
	}
}

// TestJWT_MissingToken checks that the middleware returns an error when the token is missing
func TestJWT_MissingToken(t *testing.T) {
	// Arrange
	secret := "test-secret"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	expectedResponse := map[string]string{"error": "an authorization header is required"}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{})

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

// TestJWT_MalformedToken checks that the middleware returns an error when the token is not properly formated
func TestJWT_MalformedToken(t *testing.T) {
	// Arrange
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

// TestJWT_InvalidSigninMethod checks that the middleware returns an error when the token´s signature is not expected
func TestJWT_InvalidSigninMethod(t *testing.T) {
	// Arrange
	jwtOk := "Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.e30."
	secret := "none signing method allowed"
	headerName := "Authorization"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(headerName, jwtOk)
	expectedResponse := map[string]string{"error": "invalid token: signin method not valid"}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{})

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

// TestJWT_InvalidSecret checks that the middleware returns an error when the token´s secret is not the expected one
func TestJWT_InvalidSecret(t *testing.T) {
	// Arrange
	jwtInvalidSecret := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.df4nTfuWWdndfrlIxF0iWUrrcANrM4bzKdbYa9VeAj8"
	secret := "test-secret"
	headerName := "Authorization"
	url := "http://testing"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(headerName, jwtInvalidSecret)
	expectedResponse := map[string]string{"error": "invalid token: signature is invalid"}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := JWT(handlerFunc, secret, jwt.MapClaims{})

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

// TestJWT_MissingClaim checks that the middleware returns an error when a required claim is missing in the token
func TestJWT_MissingClaim(t *testing.T) {
	// Arrange
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
