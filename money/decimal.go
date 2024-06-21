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
	// ErrPrecisionDecrease is returned when an attempt is made to decrease decimal precision which would result in errors.
	ErrPrecisionDecrease = Error("cannot decrease the precision of a decimal")
)

// Decimal represents a decimal number which can store a floating point value.
// Example: 123.45 = {units: 12345, precision: 2} (12345 * 10^-2 = 123.45)
type Decimal struct {
	// units is the integer representation of the number. Multiply this by 10^-precision to get the decimal value.
	units int64
	// precision is the number of decimal places. This is the power of 10 to multiply the units by.
	precision uint8
}

// ParseDecimal parses a string representation of a decimal number and returns a Decimal.
// The input string should be in the format "123.45" where the decimal point is optional.
// It assumes no more than a single decimal point.
func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	// maxDecimal is the number of digits in one trillion (10^12).
	const maxDecimal = 12

	if len(intPart) > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	units, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	precision := uint8(len(fracPart))
	decimal := &Decimal{units: units, precision: precision}
	decimal.simplify()
	return *decimal, nil
}

// String implements the Stringer interface.
func (d *Decimal) String() string {
	if d.precision == 0 {
		return fmt.Sprintf("%d", d.units)
	}

	divisor := pow10(d.precision)
	frac := d.units % divisor
	whole := d.units / divisor
	decimalFormat := "%d.%0" + strconv.Itoa(int(d.precision)) + "d"
	return fmt.Sprintf(decimalFormat, whole, frac)
}

// simplify removes trailing zeroes when it would not affect the value.
func (d *Decimal) simplify() {
	for d.precision > 0 && d.units%10 == 0 {
		d.units /= 10
		d.precision--
	}
}

// pow10 returns the representation of 10^power (e.g. 10^0 = 1, 10^1 = 10)
// optimized for powers of 9 or less.
func pow10(power uint8) int64 {
	values := map[uint8]int64{
		0: 1,
		1: 10,
		2: 100,
		3: 1_000,
		4: 10_000,
		5: 100_000,
		6: 1_000_000,
		7: 10_000_000,
		8: 100_000_000,
		9: 1_000_000_000,
	}

	pow, found := values[power]

	if !found {
		pow = values[9]
		for i := power - 9; i > 0; i-- {
			pow *= 10
		}
	}

	return pow
}

// updatePrecision adds additional precision to the Decimal, updating the precision and units correctly.
func (d *Decimal) updatePrecision(precision uint8) error {
	if precision < d.precision {
		return ErrPrecisionDecrease
	}

	increase := precision - d.precision

	if increase == 0 {
		return nil
	}

	d.units *= pow10(increase)
	d.precision = precision

	return nil
}
