package ecbank

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/th3oth3rjak3/MoneyConverter/money"
)

// the ecbank api compares all values to the Euro.
const baseCurrencyCode = "EUR"

const (
	errEnvelopeMissingSource = ecbankError("envelope missing source")
	errEnvelopeMissingTarget = ecbankError("envelope missing target")
	ErrUnexpectedFormat      = ecbankError("response body was not in the expected format")
	ErrExchangeRateNotFound  = ecbankError("exchange rate not found")
)

// envelope is a structure used to model the current bank api.
type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

// currencyRate contains the currency code and exchange rate.
type currencyRate struct {
	Currency string             `xml:"currency,attr"`
	Rate     money.ExchangeRate `xml:"rate,attr"`
}

// exchangeRates converts the envelope exchanges rates into a map by currency name.
func (e *envelope) exchangeRates() map[string]money.ExchangeRate {
	rates := make(map[string]money.ExchangeRate, len(e.Rates)+1)

	for _, c := range e.Rates {
		rates[c.Currency] = c.Rate
	}

	rates[baseCurrencyCode] = 1.

	return rates
}

// exchangeRate calculates the exchange rate from the source to target currency
func (e *envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		return 1., nil
	}

	rates := e.exchangeRates()

	sourceFactor, sourceFound := rates[source]
	if !sourceFound {
		return 0, errEnvelopeMissingSource
	}

	targetFactor, targetFound := rates[target]
	if !targetFound {
		return 0, errEnvelopeMissingTarget
	}

	// note: 1 / (EUR -> USD) == USD -> EUR
	// scenario: if going from USD -> CAD
	// 1 / (EUR -> USD) * EUR -> CAD == (EUR -> CAD) / (EUR -> USD) == target / source
	return targetFactor / sourceFactor, nil
}

// readRateFromResponse parses the response body and gets the exchange from the source to the target.
func readRateFromResponse(source, target string, respBody io.Reader) (money.ExchangeRate, error) {
	decoder := xml.NewDecoder(respBody)

	var ecbMessage envelope
	err := decoder.Decode(&ecbMessage)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrUnexpectedFormat, err)
	}

	rate, err := ecbMessage.exchangeRate(source, target)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrExchangeRateNotFound, err)
	}

	return rate, nil
}
