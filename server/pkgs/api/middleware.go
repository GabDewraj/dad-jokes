package api

import (
	"fmt"
	"log"
	"net/http"

	"api-keys/pkgs/secrets"

	"github.com/gorilla/mux"
)

const apiKeyHeader = "X-Dad-Jokes-Access-Token"

// func AuthAPIKey(secretId string) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		key := ctx.Request.Header.Get(apiKeyHeader)

// 		// Compare secret with key.
// 		secret, err := secrets.VerifySecret()
// 		if err != nil {
// 			log.Println("failed to get secret")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"code":    http.StatusUnauthorized,
// 				"message": http.StatusText(http.StatusUnauthorized),
// 			})
// 		}

// 		if secret != key {
// 			log.Println("key doesnt match!")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"code":    http.StatusUnauthorized,
// 				"message": http.StatusText(http.StatusUnauthorized),
// 			})
// 		}
// 		log.Println("no error during check")
// 		ctx.Next()
// 	}
// }

// Type 1
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside filterContentType middleware, before the request is handled")

		// filter requests by mime type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. JSON type is expected"))
			return
		}

		// handle the request
		handler.ServeHTTP(w, r)

		log.Println("Inside filterContentType middleware, after the request was handled")
	})
}

// Type 2
func ApiKeyAuth() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			fmt.Print("Hello from inside the middleware \n")
			apiHeader := req.Header.Get(apiKeyHeader)
			err := secrets.VerifySecret(apiHeader)
			if err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			}
			fmt.Print("Well done, you are authenticated here is your key \n")
			fmt.Printf("This is the api header %s \n", apiHeader)
			next.ServeHTTP(res, req)
		})
	}
}
