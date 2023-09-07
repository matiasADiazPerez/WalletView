package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"walletview/config"
	"walletview/internal/models"
	"walletview/internal/utils"
)

type AnkrClient struct {
}

func NewAnkrClient() AnkrClient {
	return AnkrClient{}
}

func call[T any](ctx context.Context, method string, payloadByte []byte) (models.AnkrResponse[T], error) {
	var response models.AnkrResponse[T]
	url := fmt.Sprintf("%s/?%s=", config.ANKR_URL, method)
	payload := bytes.NewReader(payloadByte)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return response, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	rawResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, err
	}
	defer rawResponse.Body.Close()
	decoder := json.NewDecoder(rawResponse.Body)
	err = decoder.Decode(&response)
	return response, err
}

func (a AnkrClient) GetTokenBallance(ctx context.Context, params models.WalletParams) ([]models.TokenBalance, models.ErrorWrapper) {
	method := "ankr_getAccountBalance"
	payloadByte, err := json.Marshal(models.Ankrbody{
		Id:      1,
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		erroW := utils.NewErrorWrapper(config.CLIENT_WALLET_ERR, 0, err)
		return []models.TokenBalance{}, erroW
	}

	resp, err := call[models.AnkrBalanceResponse](ctx, method, payloadByte)
	if err != nil {
		erroW := utils.NewErrorWrapper(config.CLIENT_WALLET_ERR, 0, err)
		return []models.TokenBalance{}, erroW
	}
	if len(resp.Result.Assets) == 0 {
		erroW := utils.NewErrorWrapper(config.CLIENT_WALLET_ERR, http.StatusNotFound, fmt.Errorf("This address doesn't have funds"))
		return []models.TokenBalance{}, erroW

	}
	return resp.Result.Assets, models.ErrorWrapper{}
}

func (a AnkrClient) GetTokensSymbols() (map[string]string, error) {
	db := make(map[string]string)
	method := "ankr_getCurrencies"
	ctx := context.Background()
	payloadByte, err := json.Marshal(models.Ankrbody{
		Id:      1,
		Jsonrpc: "2.0",
		Method:  method,
		Params: models.CurrenciesParams{
			Blockchain: "eth",
		},
	})
	if err != nil {
		return db, err
	}
	resp, err := call[models.AnkrCurrenciesResponse](ctx, method, payloadByte)
	if err != nil {
		return db, err
	}
	for _, token := range resp.Result.Currencies {
		db[token.Symbol] = token.Address
	}
	return db, nil
}

func (a AnkrClient) GetTokenPrice(ctx context.Context, params models.TokenPriceParams) (string, models.ErrorWrapper) {
	method := "ankr_getTokenPrice"
	payloadByte, err := json.Marshal(models.Ankrbody{
		Id:      1,
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		erroW := utils.NewErrorWrapper(config.CLIENT_PRICE_ERR, 0, err)
		return "", erroW
	}
	resp, err := call[models.AnkrPriceResponse](ctx, method, payloadByte)
	if err != nil {
		erroW := utils.NewErrorWrapper(config.CLIENT_PRICE_ERR, 0, err)
		return "", erroW
	}
	return resp.Result.UsdPrice, models.ErrorWrapper{}
}
