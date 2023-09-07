package wallet

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"walletview/config"
	"walletview/internal/models"
	"walletview/internal/utils"
)

func (wm WalletModule) getCurrencyPrice(currency string) (float64, models.ErrorWrapper) {
	ctx := context.Background()
	var tokenContract string
	if currency != "ETH" {
		tokenContract = wm.db[currency]
		if tokenContract == "" {
			errW := utils.NewErrorWrapper(config.WALLET_ERR, http.StatusNotFound, fmt.Errorf("Symbol not recognized"))
			return 0.0, errW
		}
	}
	priceStr, errW := wm.cli.GetTokenPrice(ctx, models.TokenPriceParams{
		Blockchain:      "eth",
		ContractAddress: tokenContract,
		SkipSyncCheck:   true,
	})
	if errW.Error != nil {
		return 0.0, errW
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		errW := utils.NewErrorWrapper(config.WALLET_ERR, 0, err)
		return 0.0, errW
	}
	if price == 0 {
		errW := utils.NewErrorWrapper(config.WALLET_ERR, http.StatusNotFound, fmt.Errorf("Symbol has no price"))
		return 0.0, errW
	}
	return price, models.ErrorWrapper{}
}

func (wm WalletModule) formatBalanceResponse(tokenBalances []models.TokenBalance, price float64, currency string) ([]string, models.ErrorWrapper) {
	var response []string
	var strResp string
	for _, token := range tokenBalances {
		if token.Blockchain != "eth" {
			continue
		}
		tokenBalance, err := strconv.ParseFloat(token.Balance, 64)
		if err != nil {
			errW := utils.NewErrorWrapper(config.WALLET_ERR, 0, err)
			return response, errW
		}
		symbol := token.TokenSymbol
		if symbol == "" {
			symbol = fmt.Sprintf("No symbol, contract Address: %s", token.ContractAddress)
		}
		tokenPrice, err := strconv.ParseFloat(token.TokenPrice, 64)
		if err != nil {
			errW := utils.NewErrorWrapper(config.WALLET_ERR, 0, err)
			return response, errW
		}
		tokenBalance = tokenBalance * tokenPrice
		if currency != "" {
			tokenBalance = tokenBalance / price
			strResp = fmt.Sprintf("%s: %v $%s", symbol, tokenBalance, currency)

		} else {
			strResp = fmt.Sprintf("%s: %v $USD", symbol, tokenBalance)
		}
		response = append(response, strResp)
	}
	return response, models.ErrorWrapper{}
}

func (wm *WalletModule) WalletBalance(address, currency string) ([]string, models.ErrorWrapper) {
	var price float64
	ctx := context.Background()
	params := models.WalletParams{
		SkipSyncCheck: true,
		WalletAddress: address,
	}
	tokenBalances, errW := wm.cli.GetTokenBallance(ctx, params)
	if errW.Error != nil {
		return []string{}, errW
	}
	if currency != "" {
		price, errW = wm.getCurrencyPrice(currency)
		if errW.Error != nil {
			return []string{}, errW
		}
	}
	resp, err := wm.formatBalanceResponse(tokenBalances, price, currency)
	return resp, err
}
