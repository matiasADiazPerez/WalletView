package models

type Ankrbody struct {
	Id      int    `json:"id,omitempty"`
	Jsonrpc string `json:"jsonrpc,omitempty"`
	Method  string `json:"method,omitempty"`
	Params  any    `json:"params,omitempty"`
}

type AnkrResponse[T any] struct {
	Id      int    `json:"id,omitempty"`
	Jsonrpc string `json:"jsonrpc,omitempty"`
	Result  T      `json:"result"`
}

/* Params */
type WalletParams struct {
	NativeFist      bool   `json:"nativeFirst"`
	OnlyWhiteListed bool   `json:"onlyWhiteListed"`
	PageSize        int    `json:"pageSize"`
	PageToken       string `json:"pageToken"`
	SkipSyncCheck   bool   `json:"skipSyncCheck,omitempty"`
	WalletAddress   string `json:"walletAddress,omitempty"`
}

type CurrenciesParams struct {
	Blockchain string `json:"blockchain"`
}

type TokenPriceParams struct {
	Blockchain      string `json:"blockchain"`
	ContractAddress string `json:"contractAddress"`
	SkipSyncCheck   bool   `json:"skipSyncCheck,omitempty"`
}

/* Response */
type AnkrBalanceResponse struct {
	TotalBalanceUsd string         `json:"totalBalanceUsd"`
	Assets          []TokenBalance `json:"assets"`
}

type AnkrCurrenciesResponse struct {
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
}

type AnkrPriceResponse struct {
	Blockchain      string `json:"blockchain"`
	ContractAddress string `json:"contractAddress"`
	UsdPrice        string `json:"usdPrice"`
}
