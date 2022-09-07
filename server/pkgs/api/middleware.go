package api

import (
	"net/http"

	"api-keys/pkgs/secrets"

	"github.com/gorilla/mux"
)

const apiKeyHeader = "X-Dad-Jokes-Access-Token"

func ApiKeyAuth(apiKeyHeader string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			apiHeader := req.Header.Get(apiKeyHeader)
			err := secrets.VerifySecret(apiHeader)
			if err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(res, req)
		})
	}
}
