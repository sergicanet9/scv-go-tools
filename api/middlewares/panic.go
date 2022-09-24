package middlewares

import (
	"net/http"

	"github.com/sergicanet9/scv-go-tools/v3/api/utils"
)

//PanicRecover is a middleware function to defer and return an error response in case of panic during the handler execution
func PanicRecover(next http.HandlerFunc) http.HandlerFunc {
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
				utils.ResponseError(w, r, nil, http.StatusInternalServerError, errorMessage)
			}
		}()
		next(w, r)
	})
}
