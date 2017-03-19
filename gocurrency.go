package gocurrency

import (
	"encoding/json"
	"errors"
	"fmt"
)

const AvailableCurrenciesURL string = "http://currencyconverter.kund.nu/api/availablecurrencies/"

const ConvertCurrencyURL string = "http://currencyconverter.kund.nu/api/?from=%s&from_amount=%f&to=%s"

type Currency struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func NewCurrency(currency string) Currency {
	return Currency{ID: currency}
}

func AvailableCurrencies() ([]Currency, error) {
	resp, err := NewRequest().Get(AvailableCurrenciesURL)
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

type ConvertCurrencyResponse struct {
	From       string  `json:"from"`
	To         string  `json:"to"`
	FromAmount float64 `json:"from_amount"`
	ToAmount   float64 `json:"to_amount"`
	Error      string  `json:"error"`
}

func ConvertCurrency(from, to Currency, amount float64) (float64, error) {
	endpoint := fmt.Sprintf(ConvertCurrencyURL, from.ID, amount, to.ID)
	resp, err := NewRequest().Get(endpoint)
	if err != nil {
		return 0, err
	}

	ccResp := ConvertCurrencyResponse{}
	err = json.Unmarshal(resp, &ccResp)
	if err != nil {
		return 0, err
	}

	if ccResp.Error != "" {
		return 0, errors.New(ccResp.Error)
	}

	return ccResp.ToAmount, nil
}
