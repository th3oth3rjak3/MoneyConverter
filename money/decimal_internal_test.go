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
