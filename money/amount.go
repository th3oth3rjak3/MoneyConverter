package money

import "fmt"

const (
	// ErrTooPrecise is returned when the precision of the input is greater than that of the currency.
	ErrTooPrecise = Error("quantity is too precise")
)

// Amount represents a quantity of money in a specific currency.
type Amount struct {
	quantity Decimal
	currency Currency
}

// NewAmount creates a new Amount with the required decimal and currency.
// the precision of quantity is expected to be less than or equal to the precision of the currency.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	// Should not have an error since we're checking for precision issues
	// in the upper guard clause.
	_ = quantity.updatePrecision(currency.precision)

	return Amount{quantity: quantity, currency: currency}, nil
}

// String implements the Stringer interface for Amount.
func (a *Amount) String() string {
	return fmt.Sprintf("%s %s", &a.quantity, a.currency)
}
