package wallet

import (
	"context"
	"walletview/internal/client"
	"walletview/internal/models"
)

func walletBalance(address, currency string) ([]models.TokenBalance, models.ErrorWrapper) {
	ctx := context.Background()
	params := models.WalletParams{
		SkipSyncCheck: true,
		WalletAddress: address,
	}
	resp, err := client.GetTokenBallance(ctx, params)
	if err.Error != nil {
		return []models.TokenBalance{}, err
	}
	return resp, models.ErrorWrapper{}
}
