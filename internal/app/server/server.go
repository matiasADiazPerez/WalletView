package server

import (
	"log"
	"net/http"
	"walletview/internal/app/wallet"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Start(wm *wallet.WalletModule) {
	router := chi.NewRouter()
	initRoutes(router, wm)
	log.Println("Start server in port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func initRoutes(router *chi.Mux, wm *wallet.WalletModule) {
	router.Get("/walletBalance", wm.HandleWalletView)
	router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))
}
