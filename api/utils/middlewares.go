package utils

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

//JWTMiddleware is a middleware function to check the authorization JWT Bearer token header of the request
func JWTMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("authorization")
		if bearerToken != "" {
			token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte(secret), nil
			})
			if err != nil {
				ResponseError(w, r, http.StatusUnauthorized, err.Error())
				return
			}
			if token.Valid {
				context.Set(r, "decoded", token.Claims)
				next.ServeHTTP(w, r)
			} else {
				ResponseError(w, r, http.StatusUnauthorized, "invalid authorization token")
			}
		} else {
			ResponseError(w, r, http.StatusUnauthorized, "an authorization header is required")
		}
	})
}

//HandlerFuncErrorHandling is a middleware function to defer and return an error response in case of panic during the request
func HandlerFuncErrorHandling(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			errorMessage := ""
			if rec := recover(); rec != nil {
				switch t := rec.(type) {
				case error:
					errorMessage = t.Error()
				case string:
					errorMessage = t
				default:
					errorMessage = "unknown error ocurred"
				}
				ResponseError(w, r, http.StatusInternalServerError, errorMessage)
			}
		}()
		next(w, r)
	})
}
