package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/sergicanet9/scv-go-tools/v3/api/utils"
)

//JWT is a middleware function to check the authorization JWT Bearer token header of the request
func JWT(next http.Handler, secret string, claims jwt.MapClaims) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an unknown error")
					}
					return []byte(secret), nil
				})
				if err != nil {
					utils.ResponseError(w, r, nil, http.StatusUnauthorized, err.Error())
					return
				}
				if c, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					for name, value := range claims {
						if claim, ok := c[name]; !(ok && claim == value) {
							utils.ResponseError(w, r, nil, http.StatusUnauthorized, fmt.Sprintf("required claim %s not found or incorrect", name))
							return
						}
					}
					context.Set(r, "decoded", token.Claims)
					next.ServeHTTP(w, r)
				} else {
					utils.ResponseError(w, r, nil, http.StatusUnauthorized, "invalid authorization token")
				}
			} else {
				utils.ResponseError(w, r, nil, http.StatusUnauthorized, "authorization header not properly formated, should be Bearer + {token}")
			}
		} else {
			utils.ResponseError(w, r, nil, http.StatusUnauthorized, "an authorization header is required")
		}
	})
}
