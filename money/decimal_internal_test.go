package money

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	type testCase struct {
		input   string
		want    Decimal
		wantErr error
	}

	testCases := map[string]testCase{
		"Empty string": {
			input:   "",
			want:    Decimal{},
			wantErr: ErrInvalidDecimal,
		},
		"Not a Number": {
			input:   "NaN",
			want:    Decimal{},
			wantErr: ErrInvalidDecimal,
		},
		"Invalid decimal": {
			input:   "123.12.12",
			want:    Decimal{},
			wantErr: ErrInvalidDecimal,
		},
		"Too large": {
			input:   "1234567890123",
			want:    Decimal{},
			wantErr: ErrTooLarge,
		},
		"nominal usage": {
			input:   "123.45",
			want:    Decimal{units: 12345, precision: 2},
			wantErr: nil,
		},
		"no decimal digits": {
			input:   "123",
			want:    Decimal{units: 123, precision: 0},
			wantErr: nil,
		},
		"decimal only with 0 for int part": {
			input:   "0.123",
			want:    Decimal{units: 123, precision: 3},
			wantErr: nil,
		},
		"int part only with a period but no decimal part": {
			input:   "123.",
			want:    Decimal{units: 123, precision: 0},
			wantErr: nil,
		},
		"period with only decimal part, no leading 0": {
			input:   ".123",
			want:    Decimal{units: 123, precision: 3},
			wantErr: nil,
		},
		"single 0 for decimal part": {
			input:   "123.0",
			want:    Decimal{units: 123, precision: 0},
			wantErr: nil,
		},
		"multiple 0's for decimal part": {
			input:   "123.000",
			want:    Decimal{units: 123, precision: 0},
			wantErr: nil,
		},
		"trailing 0's for int part and decimal part": {
			input:   "12345000.00",
			want:    Decimal{units: 12345000, precision: 0},
			wantErr: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, gotErr := ParseDecimal(tc.input)
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("ParseDecimal(%s) got error %#v; want %#v", tc.input, gotErr, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ParseDecimal(%s) got %#v; want %#v", tc.input, got, tc.want)
			}
		})
	}
}

func decimalEqual(input Decimal, target Decimal) bool {
	if input.precision != target.precision {
		return false
	}

	if input.units != target.units {
		return false
	}

	return true
}

func TestUpdatePrecision(t *testing.T) {
	type testCase struct {
		input        *Decimal
		newPrecision uint8
		want         Decimal
		wantErr      error
	}

	testCases := map[string]testCase{
		"increase by 1": {
			input:        &Decimal{units: 123, precision: 0},
			newPrecision: 1,
			want:         Decimal{units: 1230, precision: 1},
			wantErr:      nil,
		},
		"increase by 2": {
			input:        &Decimal{units: 123, precision: 0},
			newPrecision: 2,
			want:         Decimal{units: 12300, precision: 2},
			wantErr:      nil,
		},
		"no increase": {
			input:        &Decimal{units: 123, precision: 0},
			newPrecision: 0,
			want:         Decimal{units: 123, precision: 0},
			wantErr:      nil,
		},
		"decrease should not change original": {
			input:        &Decimal{units: 123, precision: 1},
			newPrecision: 0,
			want:         Decimal{units: 123, precision: 1},
			wantErr:      ErrPrecisionDecrease,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotErr := tc.input.updatePrecision(tc.newPrecision)
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("got err: %#v, want: %#v", gotErr, tc.wantErr)
			}

			if !decimalEqual(*tc.input, tc.want) {
				t.Errorf("got: %#v, want: %#v", *tc.input, tc.want)
			}
		})
	}
}

func TestPow10(t *testing.T) {
	type testCase struct {
		power uint8
		want  int64
	}

	testCases := map[string]testCase{
		"10^0 should be 1": {
			power: 0,
			want:  1,
		},
		"10^1 should be 10": {
			power: 1,
			want:  10,
		},
		"10^2 should be 100": {
			power: 2,
			want:  100,
		},
		"10^3 should be 1_000": {
			power: 3,
			want:  1_000,
		},
		"10^4 should be 10_000": {
			power: 4,
			want:  10_000,
		},
		"10^5 should be 100_000": {
			power: 5,
			want:  100_000,
		},
		"10^6 should be 1_000_000": {
			power: 6,
			want:  1_000_000,
		},
		"10^7 should be 10_000_000": {
			power: 7,
			want:  10_000_000,
		},
		"10^8 should be 100_000_000": {
			power: 8,
			want:  100_000_000,
		},
		"10^9 should be 1_000_000_000": {
			power: 9,
			want:  1_000_000_000,
		},
		"10^10 should be 10_000_000_000": {
			power: 10,
			want:  10_000_000_000,
		},
		"10^11 should be 100_000_000_000": {
			power: 11,
			want:  100_000_000_000,
		},
		"10^12 should be 1_000_000_000_000": {
			power: 12,
			want:  1_000_000_000_000,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := pow10(tc.power)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
