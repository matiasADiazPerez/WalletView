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

type WalletParams struct {
	NativeFist      bool   `json:"nativeFirst"`
	OnlyWhiteListed bool   `json:"onlyWhiteListed"`
	PageSize        int    `json:"pageSize"`
	PageToken       string `json:"pageToken"`
	SkipSyncCheck   bool   `json:"skipSyncCheck,omitempty"`
	WalletAddress   string `json:"walletAddress,omitempty"`
}

type AnkrBalanceResponse struct {
	TotalBalanceUsd string         `json:"totalBalanceUsd"`
	Assets          []TokenBalance `json:"assets"`
}

type AnkrTokenBalance struct {
	TokenBalance
	BalanceRawInteger string `json:"balanceRawInteger"`
	BalanceUsd        string `json:"balanceUsd"`
	HolderAddress     string `json:"holderAddress"`
	Thumbnail         string `json:"thumbnail"`
	TokenDecimal      int    `json:"tokenDecimal"`
	TokenSymbol       string `json:"tokenSumbol"`
}
