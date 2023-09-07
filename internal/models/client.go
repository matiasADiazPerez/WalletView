package models

type TokenBalance struct {
	TokenSymbol     string `json:"tokenSymbol,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	Blockchain      string `json:"blockchain,omitempty"`
	Balance         string `json:"balance,omitempty"`
	TokenPrice      string `json:"tokenPrice,omitempty"`
}

type ConvertParams struct {
	Blockchain      string
	ContractAddress string
}

type TokenParams struct {
	Blockchain      string `json:"blockchain,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	SkipSyncCheck   bool   `json:"skipSyncCheck,omitempty"`
}
