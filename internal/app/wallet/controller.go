package wallet

import (
	"fmt"
	"net/http"
	"walletview/config"
	"walletview/internal/client"
	"walletview/internal/utils"
)

type WalletModule struct {
	db  map[string]string
	cli client.Client
}

func NewWalletModule(db map[string]string, cli client.Client) *WalletModule {
	return &WalletModule{
		db:  db,
		cli: cli,
	}
}

func (wm *WalletModule) HandleWalletView(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	currency := r.URL.Query().Get("currency")
	if address == "" {
		err := utils.NewErrorWrapper(config.WALLET_ERR, http.StatusBadRequest, fmt.Errorf("you must give an address"))
		utils.HandleError(err, w)
		return
	}
	resp, err := wm.WalletBalance(address, currency)
	if err.Error != nil {
		utils.HandleError(err, w)
		return
	}
	utils.CreateResponse("Success", resp, w)
}
