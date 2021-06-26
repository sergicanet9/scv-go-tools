package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

// ResponseJSON makes the response with payload as json format
func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// ResponseError makes the error response with payload as json format
func ResponseError(w http.ResponseWriter, status int, message string) {
	ResponseJSON(w, status, map[string]string{"error": message})
}

//JWTMiddleware is a middleware function to check the authorization JWT Bearer token in header of requestt
func JWTMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}
					return []byte(secret), nil
				})
				if err != nil {
					ResponseError(w, http.StatusUnauthorized, err.Error())
					return
				}
				if token.Valid {
					context.Set(r, "decoded", token.Claims)
					next.ServeHTTP(w, r)
				} else {
					ResponseError(w, http.StatusUnauthorized, "Invalid authorization token")
				}
			}
		} else {
			ResponseError(w, http.StatusUnauthorized, "An authorization header is required")
		}
	})
}

//HandlerFuncErrorHandling is a middleware function to defer and return an error response in case of panic during the request
func HandlerFuncErrorHandling(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			errorMessage := ""
			if r := recover(); r != nil {
				switch t := r.(type) {
				case error:
					errorMessage = t.Error()
				case string:
					errorMessage = t
				default:
					errorMessage = "unknown error ocurred"
				}
				ResponseError(w, http.StatusInternalServerError, errorMessage)
			}
		}()
		next(w, r)
	})
}
