package server

import (
	"log"
	"net/http"
	"walletview/internal/app/auth"
	"walletview/internal/app/wallet"

	"github.com/go-chi/chi/v5"
)

func Start(wm *wallet.WalletModule, authHandler *auth.AuthHandler) {
	router := chi.NewRouter()
	initRoutes(router, wm, authHandler)
	log.Println("Start server in port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func initRoutes(router *chi.Mux, wm *wallet.WalletModule, authHandler *auth.AuthHandler) {
	router.With(authHandler.APIKeyMiddleware()).Group(func(router chi.Router) {
		router.Get("/walletBalance", wm.HandleWalletView)
	})
}
