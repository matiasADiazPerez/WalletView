package main

import (
	"log"
	"walletview/internal/app/auth"
	"walletview/internal/app/server"
	"walletview/internal/app/wallet"
	"walletview/internal/client"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cli := client.NewAnkrClient()
	db, err := cli.GetTokensSymbols()
	if err != nil {
		panic(err)
	}
	walletModule := wallet.NewWalletModule(db, cli)
	authHandler := auth.NewAuthHandler()
	server.Start(walletModule, authHandler)
}
