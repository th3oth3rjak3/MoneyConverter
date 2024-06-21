package money_test

import (
	"testing"

	"github.com/th3oth3rjak3/MoneyConverter/money"
)

// mustParseDecimal ensures that testing code correctly parses the decimal.
func mustParseDecimal(t *testing.T, value string) money.Decimal {
	t.Helper()

	dec, err := money.ParseDecimal(value)
	if err != nil {
		t.Fatalf("got error when parsing decimal: %s", err.Error())
	}

	return dec
}

func TestDecimalString(t *testing.T) {
	type testCase struct {
		input money.Decimal
		want  string
	}

	testCases := map[string]testCase{
		"decimal with no precision": {
			input: mustParseDecimal(t, "1"),
			want:  "1",
		},
		"decimal with 1 decimal place": {
			input: mustParseDecimal(t, "1.1"),
			want:  "1.1",
		},
		"decimal with 2 decimal places": {
			input: mustParseDecimal(t, "1.11"),
			want:  "1.11",
		},
		"ends with a zero": {
			input: mustParseDecimal(t, "1.10"),
			want:  "1.1",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.input.String()
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
