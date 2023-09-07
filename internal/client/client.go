package client

import (
	"context"
	"walletview/internal/models"
)

type Client interface {
	GetTokenBallance(ctx context.Context, params models.WalletParams) ([]models.TokenBalance, models.ErrorWrapper)
	GetTokensSymbols() (map[string]string, error)
	GetTokenPrice(ctx context.Context, params models.TokenPriceParams) (string, models.ErrorWrapper)
}
