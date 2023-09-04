package server

import (
	"log"
	"net/http"
	"walletview/internal/app/wallet"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Start() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	initRoutes(router)
	log.Println("Start server in port :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func initRoutes(router *chi.Mux) {
	router.Get("/walletBalance", wallet.HandleWalletView)
	router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))
}
