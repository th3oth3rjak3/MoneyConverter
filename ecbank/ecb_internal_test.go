package ecbank

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/th3oth3rjak3/MoneyConverter/money"
)

func mustParseCurrency(t *testing.T, currency string) money.Currency {
	t.Helper()

	curr, err := money.ParseCurrency(currency)
	if err != nil {
		t.Fatalf("could not parse currency: %s", err.Error())
	}

	return curr
}

func TestEuroCentralBank_FetchExchangeRate_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(
			w,
			`<?xml version="1.0" encoding="UTF-8"?>
			<gesmes:Envelope>
				<Cube>
					<Cube>
						<Cube currency="USD" rate="1.0688" />
						<Cube currency="CAD" rate="1.4632" />
					</Cube>
				</Cube>
			</gesmes:Envelope>`)
	}))

	defer ts.Close()

	ecb := EuropeanCentralBank{
		url: ts.URL,
	}

	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t, "CAD"))
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	// target / source is the same as 1 / EUR -> source * EUR -> target
	want := money.ExchangeRate(1.4632 / 1.0688)

	if got != want {
		t.Errorf("got: %#v, want: %#v", got, want)
	}
}

func TestEuroCentralBank_FetchExchangeRate_ServerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	defer ts.Close()

	ecb := EuropeanCentralBank{
		url: ts.URL,
	}

	_, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t, "CAD"))
	if !errors.Is(err, ErrServerSide) {
		t.Errorf("got: %s, want: %s", err.Error(), ErrServerSide.Error())
	}
}

func TestEuroCentralBank_FetchExchangeRate_ClientError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	defer ts.Close()

	ecb := EuropeanCentralBank{
		url: ts.URL,
	}

	_, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t, "CAD"))
	if !errors.Is(err, ErrClientSide) {
		t.Errorf("got: %s, want: %s", err.Error(), ErrClientSide.Error())
	}
}
