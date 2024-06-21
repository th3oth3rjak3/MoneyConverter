package money

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewAmount(t *testing.T) {
	type testCase struct {
		decimal  Decimal
		currency Currency
		want     Amount
		wantErr  error
	}

	testCases := map[string]testCase{
		"decimal and currency with equal precision": {
			decimal:  Decimal{units: 12345, precision: 2},
			currency: Currency{code: "USD", precision: 2},
			want: Amount{
				quantity: Decimal{units: 12345, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
			wantErr: nil,
		},
		"decimal with greater precision": {
			decimal:  Decimal{units: 12345, precision: 2},
			currency: Currency{code: "ABC", precision: 0},
			want:     Amount{},
			wantErr:  ErrTooPrecise,
		},
		"decimal with smaller precision": {
			decimal:  Decimal{units: 123, precision: 0},
			currency: Currency{code: "USD", precision: 2},
			want: Amount{
				quantity: Decimal{units: 12300, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
			wantErr: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := NewAmount(tc.decimal, tc.currency)
			if !errors.Is(tc.wantErr, err) {
				t.Errorf("got err: %#v, want: %#v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %#v, want: %#v", got, tc.want)
			}
		})
	}
}
