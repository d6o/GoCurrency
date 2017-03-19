# ![GoCurrency](http://image.prntscr.com/image/e1498265eef040bc956e388c35d8f93a.png)

# GoCurrency ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/GoCurrency) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

GoCurrency is a Go program made on the top of Kund Nu Currency Converter API.

The GoCurrency's goal is to be a perfect tool providing a stupidly easy-to-use and fast program to convert values between currencies.

[![asciicast](https://asciinema.org/a/107878.png)](https://asciinema.org/a/107878?t=10)

**Table of Contents**

- [Project Status](#project-status)
- [Features](#features)
- [Available Currencies](#available-currencies)
- [Installation](#installation)
- [Usage](#usage)
  - [Get all available currencies](#get-all-available-currencies)
  - [Convert 100 USD to all currencies](#convert-100-USD-to-all-currencies)
- [Contributing](#contributing)
  - [Bug Reports & Feature Requests](#bug-reports--feature-requests)
  - [Developing](#developing)
- [Social Coding](#social-coding)
- [License](#license)

## Project Status

GoCurrency is on beta. Pull Requests [are welcome](https://github.com/DiSiqueira/GoCurrency#social-coding)

## Features

- Open source - You can check out our code
- Secure
- Always up-to-date
- 100% satisfaction guaranteed
- It's perfect to convert amounts between currencies
- STUPIDLY [EASY TO USE](https://github.com/DiSiqueira/GoCurrency#usage)
- Very fast start up and response time
- Uses native libs

## Available Currencies

- ATS Austria, shilling
- AUD Australian, dollar
- BEF Belgien, franc
- BRL Brazilien, real
- CAD Canada, dollar
- CHF Switzerland, francs
- CNY China, yuan renminbi
- CYP Cyprus, pound
- CZK Czech Republic, koruna
- DEM Germany, mark
- DKK Denmark, krone
- EEK Estonian, kroon
- ESP Spain, pesetas
- EUR Euroland, euro
- FIM Finland, marka
- FRF France, franc
- GBP Great Britain, pound
- GRD Greece, drachmer
- HKD Hong Kong, dollar
- HUF Hungary, forint
- IDR Indonesia, rupiah
- IEP Ireland, pund
- INR India, rupee
- ISK Iceland, kronor
- ITL Italy, lire
- JPY Japan, yen
- KRW South Korea, won
- KWD Kuwait, dinar
- LTL Lithuania,  litas
- LVL Latvia, lat
- MAD Morocko, dirham
- MXN Mexico, nuevo peso
- MYR Malaysia, ringgit
- NLG Dutchland, guilder
- NOK Norway, krone
- NZD New Zealand, dollar
- PLN Poland, zloty
- PTE Portugal, escudo
- RUB Russia, rouble
- SAR Saudi Arabia, riyal
- SEK Sweden, kronor
- SGD Singapore, dollar
- SIT Slovenia, tolar
- SKK Slovakia, koruna
- THB Thailand, baht
- TRL Turkey, lira
- TRY Turkey, new lira
- USD US, dollar
- ZAR South Africa, rand

## Installation

### Option 1: Go Get

```bash
$ go get github.com/disiqueira/gocurrency
```

## Usage

### Get all available currencies

```golang
package main

import (
	"fmt"

	"github.com/disiqueira/gocurrency"
)

func main() {
	curList, err := gocurrency.AvailableCurrencies()

	for _, currency := range curList {
		fmt.Println(currency.Description)
	}
}
```

![](http://image.prntscr.com/image/996b6db6daa5404daa52b551849da8f3.png)

### Convert 100 USD to all currencies

```golang
package main

import (
	"fmt"
        "strconv"

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
```

![](http://image.prntscr.com/image/4aac591db36443698437e0d60a63fb88.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/GoCurrency/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ git clone https://github.com/disiqueira/gocurrency.git gocurrency
$ cd gocurrency/
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/GoCurrency/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
