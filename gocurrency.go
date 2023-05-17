package gocurrency

import (
	"encoding/json"
	"fmt"
	"strings"
)

const availableCurrenciesURL = "http://api.fixer.io/"

// GoCurrency contains rates for all currencies
type GoCurrency struct {
	base  string
	date  string
	rates map[string]interface{}
}

// Parameters are passed in the exported function
type Parameters struct {
	base    string
	date    string
	symbols []string
}

// GetCurrencies returns a GoCurrency struct
func GetCurrencies(pms *Parameters) (*GoCurrency, error) {
	// setting up parameters
	getParams := make([]string, 0)
	if pms.date == "" {
		pms.date = "latest"
	}
	if pms.base != "" {
		getParams = append(getParams, fmt.Sprintf("base=%s", pms.base))
	}
	if len(pms.symbols) > 0 {
		getParams = append(getParams, fmt.Sprintf("symbols=%s", strings.Join(pms.symbols, ",")))
	}
	URL := fmt.Sprintf("%s%s", availableCurrenciesURL, pms.date)
	if len(getParams) > 0 {
		URL += "?" + strings.Join(getParams, "&")
	}
	// actual request
	rawResp, err := newRequest().Get(URL)
	if err != nil {
		return nil, err
	}
	// response handling
	var resp map[string]interface{}
	err = json.Unmarshal(rawResp, &resp)
	if err != nil {
		return nil, err
	}
	goCurrency := &GoCurrency{
		base:  resp["base"].(string),
		date:  resp["date"].(string),
		rates: resp["rates"].(map[string]interface{}),
	}
	return goCurrency, nil
}
