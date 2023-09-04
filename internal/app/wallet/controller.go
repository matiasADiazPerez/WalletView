package wallet

import (
	"fmt"
	"net/http"
	"walletview/config"
	"walletview/internal/utils"
)

func HandleWalletView(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	currency := r.URL.Query().Get("currency")
	if address == "" {
		err := utils.NewErrorWrapper(config.WALLET_ERR, http.StatusBadRequest, fmt.Errorf("you must give an address"))
		utils.HandleError(err, w)
	}
	if currency == "" {
		currency = "eth"
	}

	resp, err := walletBalance(address, currency)
	if err.Error != nil {
		utils.HandleError(err, w)
	}
	utils.CreateResponse("Success", resp, w)
}
