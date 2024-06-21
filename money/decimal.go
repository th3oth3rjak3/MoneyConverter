package money

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ErrInvalidDecimal is returned when the input is malformed.
	ErrInvalidDecimal = Error("unable to convert the decimal")
	// ErrTooLarge is returned when the input is too large which would cause floating point precision errors.
	ErrTooLarge = Error("quantity over 10^12 is too large")
)

// Decimal represents a decimal number which can store a floating point value.
// Example: 123.45 = {units: 12345, precision: 2} (12345 * 10^-2 = 123.45)
type Decimal struct {
	// units is the integer representation of the number. Multiply this by 10^-precision to get the decimal value.
	units int64
	// precision is the number of decimal places. This is the power of 10 to multiply the units by.
	precision byte
}

// ParseDecimal parses a string representation of a decimal number and returns a Decimal.
// The input string should be in the format "123.45" where the decimal point is optional.
// It assumes no more than a single decimal point.
func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	// maxDecimal is the number of digits in one trillion.
	const maxDecimal = 12

	if len(intPart) > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	units, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	precision := byte(len(fracPart))
	decimal := &Decimal{units: units, precision: precision}
	decimal.simplify()
	return *decimal, nil
}

// simplify removes trailing zeroes when it would not affect the value.
func (d *Decimal) simplify() {
	for d.precision > 0 && d.units%10 == 0 {
		d.units /= 10
		d.precision--
	}
}
