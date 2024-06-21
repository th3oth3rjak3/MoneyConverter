package money

import (
	"reflect"
	"testing"
)

func TestApplyExchangeRate(t *testing.T) {
	type testCase struct {
		in       Amount
		rate     ExchangeRate
		currency Currency
		want     Amount
	}

	testCases := map[string]testCase{
		"rate of 1.00": {
			in: Amount{
				quantity: Decimal{units: 12300, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
			rate:     ExchangeRate(Decimal{units: 100, precision: 2}),
			currency: Currency{code: "TST", precision: 2},
			want: Amount{
				quantity: Decimal{units: 12300, precision: 2},
				currency: Currency{code: "TST", precision: 2},
			},
		},
		"larger converted precision": {
			in: Amount{
				quantity: Decimal{units: 12300, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
			rate:     ExchangeRate(Decimal{units: 11111, precision: 4}),
			currency: Currency{code: "TST", precision: 2},
			want: Amount{
				quantity: Decimal{units: 13666, precision: 2},
				currency: Currency{code: "TST", precision: 2},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := applyExchangeRate(tc.in, tc.currency, tc.rate)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %#v, want: %#v", got, tc.want)
			}
		})
	}
}
