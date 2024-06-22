package ecbank

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/th3oth3rjak3/MoneyConverter/money"
)

func TestEnvelopeExchangeRates(t *testing.T) {
	type testCase struct {
		envelope *envelope
		want     map[string]money.ExchangeRate
	}

	testCases := map[string]testCase{
		"empty list": {
			envelope: &envelope{
				Rates: []currencyRate{},
			},
			want: map[string]money.ExchangeRate{
				"EUR": 1.,
			},
		},
		"some values": {
			envelope: &envelope{
				Rates: []currencyRate{
					{
						Currency: "USD",
						Rate:     1.1,
					},
					{
						Currency: "CAD",
						Rate:     1.3,
					},
				},
			},
			want: map[string]money.ExchangeRate{
				"USD": 1.1,
				"CAD": 1.3,
				"EUR": 1.0,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.envelope.exchangeRates()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %#v, want: %#v", got, tc.want)
			}
		})
	}
}

func TestEnvelopeExchangeRate(t *testing.T) {
	type testCase struct {
		from     string
		to       string
		envelope *envelope
		want     money.ExchangeRate
		wantErr  error
	}

	testCases := map[string]testCase{
		"to same currency": {
			from: "USD",
			to:   "USD",
			envelope: &envelope{
				Rates: []currencyRate{
					{
						Currency: "USD",
						Rate:     1.2,
					},
				},
			},
			want:    1.0,
			wantErr: nil,
		},
		"USD to CAD": {
			from: "USD",
			to:   "CAD",
			envelope: &envelope{
				Rates: []currencyRate{
					{
						Currency: "USD",
						Rate:     1.4,
					},
					{
						Currency: "CAD",
						Rate:     1.2,
					},
				},
			},
			want:    0.857142857,
			wantErr: nil,
		},
		"missing source": {
			from: "USD",
			to:   "CAD",
			envelope: &envelope{
				Rates: []currencyRate{
					{
						Currency: "CAD",
						Rate:     1.2,
					},
				},
			},
			want:    0,
			wantErr: errEnvelopeMissingSource,
		},
		"missing target": {
			from: "USD",
			to:   "CAD",
			envelope: &envelope{
				Rates: []currencyRate{
					{
						Currency: "USD",
						Rate:     1.4,
					},
				},
			},
			want:    0,
			wantErr: errEnvelopeMissingTarget,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := tc.envelope.exchangeRate(tc.from, tc.to)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("got error: %s, wanted error: %s", err, tc.wantErr)
			}
			if !almostEqual(float64(got), float64(tc.want)) {
				t.Errorf("got: %g, want: %g", got, tc.want)
			}
		})
	}
}

// almostEqual compares two floats to 9 decimal places
func almostEqual(f1, f2 float64) bool {
	const threshold float64 = 1e-9
	return math.Abs(f1-f2) < threshold
}
