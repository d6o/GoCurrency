package gocurrency

import (
	"testing"
)

func TestAvailableCurrenciesBases(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"EUR", "EUR"},
		{"HKD", "HKD"},
		{"", "EUR"},
	}
	for _, c := range cases {
		goCurrency, _ := GetCurrencies(&Parameters{
			base: c.in,
		})
		if goCurrency.base != c.want {
			t.Errorf("GetCurrencies(%q) == %q, want %q", c.in, goCurrency.base, c.want)
		}
	}
}

func TestAvailableCurrenciesDates(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"2013-02-04", "2013-02-04"},
		{"2012-01-02", "2012-01-02"},
		{"", "latest"},
	}
	for _, c := range cases {
		goCurrency, _ := GetCurrencies(&Parameters{
			date: c.in,
		})
		if goCurrency.date != c.want && c.in != "" {
			t.Errorf("GetCurrencies(%q) == %q, want %q", c.in, goCurrency.date, c.want)
		}
	}
}

func TestAvailableCurrenciesSymbols(t *testing.T) {
	cases := []struct{
		base string
		symbols []string
		numRates int
	}{
		{"", []string{"CAD", "ZAR", "RUB"}, 3},
		{"EUR", []string{"SEK"}, 1},
		{"DKK", []string{"EUR"}, 1},
		{"USD", []string{}, 31},
	}
	for _, c := range cases {
		goCurrency, _ := GetCurrencies(&Parameters{
			base: c.base,
			symbols: c.symbols,
		})
		if c.numRates != len(goCurrency.rates) {
			t.Errorf("numRates == %q, want %q", len(goCurrency.rates), c.numRates)
		}
		for _, symbol := range c.symbols {
			if _, ok := goCurrency.rates[symbol]; ok != true {
				t.Errorf("Symbol %q not found in map", symbol, goCurrency.rates)
			}
		}
	}
}
