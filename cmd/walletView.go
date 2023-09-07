package main

import (
	"walletview/internal/app/server"
	"walletview/internal/app/wallet"
	"walletview/internal/client"
)

func main() {
	cli := client.NewAnkrClient()
	db, err := cli.GetTokensSymbols()
	if err != nil {
		panic(err)
	}
	walletModule := wallet.NewWalletModule(db, cli)
	server.Start(walletModule)
}
