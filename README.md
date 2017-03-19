# ![GoCurrency](http://image.prntscr.com/image/e1498265eef040bc956e388c35d8f93a.png)

# GoCurrency ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/GoCurrency) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

GoCurrency is a Go program made on the top of Kund Nu Currency Converter API.

The GoCurrency's goal is to be a perfect tool providing a stupidly easy-to-use and fast program to convert values between currencies.

**Table of Contents**

- [Demo](#demo)
- [Project Status](#project-status)
- [Features](#features)
- [Installation](#installation)
- [Available Currencies](#available-currencies)
- [Usage](#usage)
  - [Get all available currencies](#get-all-available-currencies)
  - [Convert 100 USD to all currencies](#convert-100-USD-to-all-currencies)
- [Contributing](#contributing)
  - [Bug Reports & Feature Requests](#bug-reports--feature-requests)
  - [Developing](#developing)
- [Social Coding](#social-coding)
- [License](#license)

## Demo

[![asciicast](http://image.prntscr.com/image/2f33d4153f794d15bd95d2d533adab98.png)](https://asciinema.org/a/107878?t=10)

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

## Installation

### Option 1: Go Get

```bash
$ go get github.com/disiqueira/gocurrency
```

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

## Usage

### Get all available currencies

```golang
package main

import (
	"fmt"

	"github.com/disiqueira/gocurrency"
)

func main() {
	curList, _ := gocurrency.AvailableCurrencies()

	for _, currency := range curList {
		fmt.Println(currency.Description)
	}
}
```

Output:
```
SEK Sweden, kronor
ATS Austria, shilling
AUD Australian, dollar
BEF Belgien, franc
BRL Brazilien, real
CAD Canada, dollar
CHF Switzerland, francs
CNY China, yuan renminbi
CYP Cyprus, pound
CZK Czech Republic, koruna
DEM Germany, mark
DKK Denmark, krone
EEK Estonian, kroon
ESP Spain, pesetas
EUR Euroland, euro
FIM Finland, marka
FRF France, franc
GBP Great Britain, pound
GRD Greece, drachmer
HKD Hong Kong, dollar
HUF Hungary, forint
IDR Indonesia, rupiah
IEP Ireland, pund
INR India, rupee
ISK Iceland, kronor
ITL Italy, lire
JPY Japan, yen
KRW South Korea, won
KWD Kuwait, dinar
LTL Lithuania,  litas
LVL Latvia, lat
MAD Morocko, dirham
MXN Mexico, nuevo peso
MYR Malaysia, ringgit
NLG Dutchland, guilder
NOK Norway, krone
NZD New Zealand, dollar
PLN Poland, zloty
PTE Portugal, escudo
RUB Russia, rouble
SAR Saudi Arabia, riyal
SGD Singapore, dollar
SIT Slovenia, tolar
SKK Slovakia, koruna
THB Thailand, baht
TRL Turkey, lira
TRY Turkey, new lira
USD US, dollar
ZAR South Africa, rand
```

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

Output:
```
USD 100.00 --> SEK 881.12
USD 100.00 --> ATS 1334.36
USD 100.00 --> AUD 130.28
USD 100.00 --> BEF 3911.85
USD 100.00 --> BRL 312.07
USD 100.00 --> CAD 133.35
USD 100.00 --> CHF 99.59
USD 100.00 --> CNY 690.75
USD 100.00 --> CYP 54.42
USD 100.00 --> CZK 2509.74
USD 100.00 --> DEM 189.66
USD 100.00 --> DKK 690.54
USD 100.00 --> EEK 1531.05
USD 100.00 --> ESP 16134.77
USD 100.00 --> EUR 92.88
USD 100.00 --> FIM 576.57
USD 100.00 --> FRF 636.10
USD 100.00 --> GBP 80.95
USD 100.00 --> GRD 33042.83
USD 100.00 --> HKD 776.45
USD 100.00 --> HUF 28765.63
USD 100.00 --> IDR 1333010.59
USD 100.00 --> IEP 76.37
USD 100.00 --> INR 6553.27
USD 100.00 --> ISK 10876.95
USD 100.00 --> ITL 187751.97
USD 100.00 --> JPY 11334.92
USD 100.00 --> KRW 113239.94
USD 100.00 --> KWD 37.74
USD 100.00 --> LTL 319.77
USD 100.00 --> LVL 69.22
USD 100.00 --> MAD 999.46
USD 100.00 --> MXN 1924.26
USD 100.00 --> MYR 489.51
USD 100.00 --> NLG 213.70
USD 100.00 --> NOK 848.16
USD 100.00 --> NZD 143.12
USD 100.00 --> PLN 399.87
USD 100.00 --> PTE 19441.33
USD 100.00 --> RUB 5790.71
USD 100.00 --> SAR 375.02
USD 100.00 --> SGD 140.29
USD 100.00 --> SIT 23310.05
USD 100.00 --> SKK 0.00
USD 100.00 --> THB 3490.97
USD 100.00 --> TRL 176224000.00
USD 100.00 --> TRY 360.64
USD 100.00 --> USD 100.00
USD 100.00 --> ZAR 1277.17
```

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
