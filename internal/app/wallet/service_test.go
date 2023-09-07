package wallet

import (
	"fmt"
	"net/http"
	"testing"
	"walletview/config"
	"walletview/internal/models"
	"walletview/internal/utils"
	"walletview/mocks"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/mock"
)

var db = map[string]string{
	"BAD": "0xlinkaddrss01",
}

type wallteTestCase struct {
	address          string
	currency         string
	expectedResponse []string
	expectedErr      models.ErrorWrapper
	BalanceResponse  []models.TokenBalance
	BalanceError     models.ErrorWrapper
	PriceResponse    string
	PriceError       models.ErrorWrapper
	ExpectPrice      bool
}

func TestWalletBalance(t *testing.T) {
	testCases := []wallteTestCase{
		{
			address: "0x01e5add8e55",
			expectedResponse: []string{
				"USDC: 1000.1 $USD",
				"LINK: 200 $USD",
			},
			BalanceResponse: []models.TokenBalance{
				{
					Blockchain:  "eth",
					Balance:     "1000.1",
					TokenPrice:  "1",
					TokenSymbol: "USDC",
				},
				{
					Blockchain:  "eth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "LINK",
				},
				{
					Blockchain:  "notEth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "NETH",
				},
			},
		},
		{
			address:     "0x01e5add8e55",
			currency:    "BAD",
			ExpectPrice: true,
			expectedResponse: []string{
				"USDC: 500 $BAD",
				"LINK: 100 $BAD",
			},
			BalanceResponse: []models.TokenBalance{
				{
					Blockchain:  "eth",
					Balance:     "1000",
					TokenPrice:  "1",
					TokenSymbol: "USDC",
				},
				{
					Blockchain:  "eth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "LINK",
				},
				{
					Blockchain:  "notEth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "NETH",
				},
			},
			PriceResponse: "2.0",
		},
		{
			address:          "0x01e5add8e55",
			expectedResponse: nil,
			BalanceResponse:  []models.TokenBalance{},
		},
		{
			address:          "0x01e5add8e55",
			expectedResponse: []string{},
			BalanceResponse:  []models.TokenBalance{},
			BalanceError:     utils.NewErrorWrapper("cli error", 0, fmt.Errorf("Something went wrong")),
			expectedErr:      utils.NewErrorWrapper("cli error", 0, fmt.Errorf("Something went wrong")),
		},
		{
			address:          "0x01e5add8e55",
			currency:         "NotBAD",
			expectedResponse: []string{},
			BalanceResponse: []models.TokenBalance{
				{
					Blockchain:  "eth",
					Balance:     "1000",
					TokenPrice:  "1",
					TokenSymbol: "USDC",
				},
				{
					Blockchain:  "eth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "LINK",
				},
				{
					Blockchain:  "notEth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "NETH",
				},
			},
			expectedErr: utils.NewErrorWrapper(config.WALLET_ERR, http.StatusNotFound, fmt.Errorf("Symbol not recognized")),
		},
		{
			address:          "0x01e5add8e55",
			currency:         "BAD",
			ExpectPrice:      true,
			expectedResponse: []string{},
			BalanceResponse: []models.TokenBalance{
				{
					Blockchain:  "eth",
					Balance:     "1000",
					TokenPrice:  "1",
					TokenSymbol: "USDC",
				},
				{
					Blockchain:  "eth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "LINK",
				},
				{
					Blockchain:  "notEth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "NETH",
				},
			},
			PriceResponse: "",
			PriceError:    utils.NewErrorWrapper(config.WALLET_ERR, 0, fmt.Errorf("Something went wrong")),
			expectedErr:   utils.NewErrorWrapper(config.WALLET_ERR, 0, fmt.Errorf("Something went wrong")),
		},
		{
			address:          "0x01e5add8e55",
			currency:         "BAD",
			ExpectPrice:      true,
			expectedResponse: []string{},
			BalanceResponse: []models.TokenBalance{
				{
					Blockchain:  "eth",
					Balance:     "1000",
					TokenPrice:  "1",
					TokenSymbol: "USDC",
				},
				{
					Blockchain:  "eth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "LINK",
				},
				{
					Blockchain:  "notEth",
					Balance:     "2",
					TokenPrice:  "100",
					TokenSymbol: "NETH",
				},
			},
			PriceResponse: "0",
			expectedErr:   utils.NewErrorWrapper(config.WALLET_ERR, http.StatusNotFound, fmt.Errorf("Symbol has no price")),
		},
	}
	t.Run("wallet view service test", func(t *testing.T) {
		for _, tc := range testCases {
			cli := mocks.NewClient(t)
			cli.On("GetTokenBallance", mock.Anything, mock.Anything).Return(tc.BalanceResponse, tc.BalanceError)
			if tc.ExpectPrice {
				cli.On("GetTokenPrice", mock.Anything, mock.Anything).Return(tc.PriceResponse, tc.PriceError)
			}
			service := NewWalletModule(db, cli)
			resp, errWrapper := service.WalletBalance(tc.address, tc.currency)

			if errWrapper.Error != nil && tc.expectedErr.Error == nil {
				t.Errorf("unexpected error: %v", errWrapper.Error.Error())
			} else if errWrapper.Error == nil && tc.expectedErr.Error != nil {
				t.Errorf("was expecting error: %v", tc.expectedErr.Error.Error())

			} else if errWrapper.Error != tc.expectedErr.Error {
				if errWrapper.Error.Error() != tc.expectedErr.Error.Error() {
					t.Errorf("got %v, want %v", errWrapper.Error, tc.expectedErr.Error)
				}

			}

			if diff := cmp.Diff(errWrapper, tc.expectedErr, cmpopts.IgnoreFields(models.ErrorWrapper{}, "Error")); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(resp, tc.expectedResponse); diff != "" {
				t.Errorf(diff)
			}

		}
	})
}
