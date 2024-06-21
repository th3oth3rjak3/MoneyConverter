package money_test

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/th3oth3rjak3/MoneyConverter/money"
// )

// // mustParseCurrency is a helper method that ensures valid currency is provided
// // to the Convert tests.
// func mustParseCurrency(t *testing.T, code string) money.Currency {
// 	t.Helper()

// 	currency, err := money.ParseCurrency(code)
// 	if err != nil {
// 		t.Fatalf("cannot parse currency code: %s", code)
// 	}

// 	return currency
// }

// // mustParseAmount is a helper method that ensures valid amounts can be created
// // for testing Convert.
// func mustParseAmount(t *testing.T, value string, code string) money.Amount {
// 	t.Helper()

// 	decimal, err := money.ParseDecimal(value)
// 	if err != nil {
// 		t.Fatalf("invalid number: %s", value)
// 	}

// 	currency, err := money.ParseCurrency(code)
// 	if err != nil {
// 		t.Fatalf("invalid currency code: %s", code)
// 	}

// 	amount, err := money.NewAmount(decimal, currency)
// 	if err != nil {
// 		t.Fatalf("cannot create new amount with value %#v and currency code %s", decimal, code)
// 	}

// 	return amount
// }

// func TestConvert(t *testing.T) {
// 	type testCase struct {
// 		amount   money.Amount
// 		to       money.Currency
// 		validate func(t *testing.T, got money.Amount, err error)
// 	}

// 	testCases := map[string]testCase{
// 		"11.22 USD to EUR": {
// 			amount: mustParseAmount(t, "11.22", "USD"),
// 			to:     mustParseCurrency(t, "EUR"),
// 			validate: func(t *testing.T, got money.Amount, err error) {
// 				if err != nil {
// 					t.Errorf("unexpected error: %v", err)
// 				}
// 				expected := money.Amount{}
// 				if !reflect.DeepEqual(got, expected) {
// 					t.Errorf("got %v, want %v", got, expected)
// 				}
// 			},
// 		},
// 	}

// 	for name, tc := range testCases {
// 		t.Run(name, func(t *testing.T) {
// 			got, err := money.Convert(tc.amount, tc.to)
// 			tc.validate(t, got, err)
// 		})
// 	}
// }
