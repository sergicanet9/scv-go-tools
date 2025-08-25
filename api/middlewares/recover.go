package middlewares

import (
	"fmt"
	"net/http"

	"github.com/sergicanet9/scv-go-tools/v3/api/utils"
)

// Recover is a middleware function to defer and return an error response in case of panic during the handler execution
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			var err error
			if rec := recover(); rec != nil {
				switch t := rec.(type) {
				case error:
					err = t
				case string:
					err = fmt.Errorf("%s", t)
				default:
					err = fmt.Errorf("unknown error ocurred")
				}
				utils.ResponseError(w, r, nil, err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
