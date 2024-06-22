package ecbank

import (
	"fmt"
	"net/http"

	"github.com/th3oth3rjak3/MoneyConverter/money"
)

const (
	// ErrCallingServer returned when an error occurs while calling the server to retrieve exchange rates.
	ErrCallingServer = ecbankError("error calling server")
	// ErrClientSide returned when a malformed client-side requests results in a 400 series error.
	ErrClientSide = ecbankError("client-side error occurred")
	// ErrServerSide returned when a server malfunction occurs and returns a 500 series error.
	ErrServerSide = ecbankError("server-side error occurred")
	// ErrUnknownStatusCode returned when any other status code is returned.
	ErrUnknownStatusCode = ecbankError("unknown status code")
)

const (
	// errors like 400 or 404
	clientErrorClass = 4
	// errors like 500
	serverErrorClass = 5
)

// EuropeanCentralBank represents a structure that can call the bank to get exchange rates.
type EuropeanCentralBank struct {
	url string
}

// FetchExchangeRate gets the exchange rate for the source to target currency.
func (ecb EuropeanCentralBank) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const ecbExchangeRateUrl string = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	if ecb.url == "" {
		ecb.url = ecbExchangeRateUrl
	}

	resp, err := http.Get(ecb.url)
	if err != nil {
		return money.ExchangeRate(0), fmt.Errorf("%w: %s", ErrCallingServer, err.Error())
	}
	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate(0), err
	}

	rate, err := readRateFromResponse(source.ISOCode(), target.ISOCode(), resp.Body)
	if err != nil {
		return 0., err
	}

	return rate, nil
}

// checkStatusCode evaluates an http status code and returns an error if not success.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w, %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w, %d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w, %d", ErrUnknownStatusCode, statusCode)
	}
}

// httpStatusClass returns the first digit of the status code.
func httpStatusClass(statusCode int) int {
	return statusCode / 100
}
