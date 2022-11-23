package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/sergicanet9/scv-go-tools/v3/api/utils"
	"github.com/sergicanet9/scv-go-tools/v3/wrappers"
)

// JWT is a middleware function to check the authorization JWT Bearer token header of the request
func JWT(next http.Handler, secret string, claims jwt.MapClaims) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("signin method not valid")
					}
					return []byte(secret), nil
				})
				if _, ok := token.Claims.(jwt.MapClaims); err != nil || !ok || !token.Valid {
					utils.ResponseError(w, r, nil, wrappers.NewUnauthorizedErr(fmt.Errorf("invalid token: %s", err.Error())))
					return
				}
				for name, value := range claims {
					if claim, ok := token.Claims.(jwt.MapClaims)[name]; !(ok && claim == value) {
						utils.ResponseError(w, r, nil, wrappers.NewUnauthorizedErr(fmt.Errorf("required claim %s not found or incorrect", name)))
						return
					}
				}
				context.Set(r, "decoded", token.Claims)
				next.ServeHTTP(w, r)
			} else {
				utils.ResponseError(w, r, nil, wrappers.NewUnauthorizedErr(fmt.Errorf("authorization header not properly formated, should be Bearer + {token}")))
			}
		} else {
			utils.ResponseError(w, r, nil, wrappers.NewUnauthorizedErr(fmt.Errorf("an authorization header is required")))
		}
	})
}
