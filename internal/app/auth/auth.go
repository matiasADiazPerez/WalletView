package auth

import (
	"fmt"
	"net/http"
	"os"
	"walletview/config"
	"walletview/internal/utils"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) APIKeyMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("x-api-key")
			if os.Getenv("API_KEY") == "" {
				utils.HandleError(utils.NewErrorWrapper(config.AUTH, http.StatusUnauthorized, fmt.Errorf("The API KEY env is not set")), w)
				return
			}
			if apiKey == "" {
				utils.HandleError(utils.NewErrorWrapper(config.AUTH, http.StatusUnauthorized, fmt.Errorf("Missing API KEY")), w)
				return
			}
			if apiKey != os.Getenv("API_KEY") {
				utils.HandleError(utils.NewErrorWrapper(config.AUTH, http.StatusUnauthorized, fmt.Errorf("Wrong API KEY")), w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
