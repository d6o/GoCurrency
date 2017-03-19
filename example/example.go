package main

import (
	"fmt"

	"github.com/disiqueira/gocurrency"
)

func main() {
	curList, _ := gocurrency.AvailableCurrencies()

	dollar := gocurrency.NewCurrency("USD")
	amount := 100.00

	for _, currency := range curList {
		conv, _ := gocurrency.ConvertCurrency(dollar, currency, amount)

		fmt.Printf("%-3s %-6.2f --> %-3s %-11.2f\n", dollar.ID, amount, currency.ID, conv)
	}
}
