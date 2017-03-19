package gocurrency

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

const availableCurrenciesURL string = "http://currencyconverter.kund.nu/api/availablecurrencies/"

const convertCurrencyURL string = "http://currencyconverter.kund.nu/api/?from=%s&from_amount=%s&to=%s"

// Currency contains ID and Description of an Currency.
type Currency struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// NewCurrency Create an instance of Currency.
func NewCurrency(currency string) Currency {
	return Currency{ID: currency}
}

// AvailableCurrencies returns an array with all currencies that can be used.
func AvailableCurrencies() ([]Currency, error) {
	resp, err := newRequest().Get(availableCurrenciesURL)
	if err != nil {
		return nil, err
	}

	curList := []Currency{}
	err = json.Unmarshal(resp, &curList)
	if err != nil {
		return nil, err
	}

	return curList, nil
}

type convertCurrencyResponse struct {
	From       string  `json:"from"`
	To         string  `json:"to"`
	FromAmount float64 `json:"from_amount"`
	ToAmount   float64 `json:"to_amount"`
	Error      string  `json:"error"`
}

// ConvertCurrency Converts an amount from one currency to another currency.
func ConvertCurrency(from, to Currency, amount decimal.Decimal) (decimal.Decimal, error) {
	endpoint := fmt.Sprintf(convertCurrencyURL, from.ID, amount, to.ID)
	resp, err := newRequest().Get(endpoint)
	if err != nil {
		return decimal.NewFromFloat(0), err
	}

	ccResp := convertCurrencyResponse{}
	err = json.Unmarshal(resp, &ccResp)
	if err != nil {
		return decimal.NewFromFloat(0), err
	}

	if ccResp.Error != "" {
		return decimal.NewFromFloat(0), errors.New(ccResp.Error)
	}

	return decimal.NewFromFloat(ccResp.ToAmount), nil
}
