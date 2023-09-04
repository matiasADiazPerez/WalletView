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

func GetTokenBallance(ctx context.Context, params models.WalletParams) ([]models.TokenBalance, models.ErrorWrapper) {
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
	return resp.Result.Assets, models.ErrorWrapper{}
}
