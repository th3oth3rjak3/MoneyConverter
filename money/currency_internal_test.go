package money

import (
	"testing"
	"errors"
)

func TestParseCurrency(t *testing.T) {
	type testCase struct {
		input   string
		want	Currency
		wantErr error
	}

	testCases := map[string]testCase{
		"Empty string": {
			input: "",
			want: Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		"Invalid currency code": {
			input: "US",
			want: Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		"incorrect symbols": {
			input: "US$",
			want: Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		"nominal usage": {
			input: "USD",
			want: Currency{code: "USD", precision: 2},
			wantErr: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.input)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("unexpected error: %v", err)
			}
			if got.code != tc.want.code || got.precision != tc.want.precision {
				t.Errorf("unexpected result: got: %v, want: %v", got, tc.want)
			}
		})
	}
}