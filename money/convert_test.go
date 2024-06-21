package money_test

import (
	"reflect"
	"testing"

	"github.com/th3oth3rjak3/MoneyConverter/money"
)

func TestConvert(t *testing.T) {
	type testCase struct {
		amount   money.Amount
		to       money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}

	testCases := map[string]testCase{
		"Empty Amount": {
			amount: money.Amount{},
			to:     money.Currency{},
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				expected := money.Amount{}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("got %v, want %v", got, expected)
				}
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to)
			tc.validate(t, got, err)
		})
	}
}
