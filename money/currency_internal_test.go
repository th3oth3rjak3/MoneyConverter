package money

import (
	"errors"
	"testing"
)

func TestParseCurrency(t *testing.T) {
	type testCase struct {
		input   string
		want    Currency
		wantErr error
	}

	testCases := map[string]testCase{
		"Empty string": {
			input:   "",
			want:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		"Invalid currency code": {
			input:   "US",
			want:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		"incorrect symbols": {
			input:   "US$",
			want:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		"United States Dollars": {
			input:   "USD",
			want:    Currency{code: "USD", precision: 2},
			wantErr: nil,
		},
		"Iranian Rial": {
			input:   "IRR",
			want:    Currency{code: "IRR", precision: 0},
			wantErr: nil,
		},
		"Chinese Yuan": {
			input:   "CNY",
			want:    Currency{code: "CNY", precision: 1},
			wantErr: nil,
		},
		"Vietnamese Dong": {
			input:   "VND",
			want:    Currency{code: "VND", precision: 1},
			wantErr: nil,
		},
		"Bahraini Dinar": {
			input:   "BHD",
			want:    Currency{code: "BHD", precision: 3},
			wantErr: nil,
		},
		"Iraqi Dinar": {
			input:   "IQD",
			want:    Currency{code: "IQD", precision: 3},
			wantErr: nil,
		},
		"Kuwaiti Dinar": {
			input:   "KWD",
			want:    Currency{code: "KWD", precision: 3},
			wantErr: nil,
		},
		"Libyan Dinar": {
			input:   "LYD",
			want:    Currency{code: "LYD", precision: 3},
			wantErr: nil,
		},
		"Omani Rial": {
			input:   "OMR",
			want:    Currency{code: "OMR", precision: 3},
			wantErr: nil,
		},
		"Tunisian Dinar": {
			input:   "TND",
			want:    Currency{code: "TND", precision: 3},
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
