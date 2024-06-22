package money

import "fmt"

// ExchangeRate represents a rate to convert from one currency to another.
type ExchangeRate float64

// Convert applies an exchange rate to convert an input amount to a target currency.
func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	exchangeRate, err := rates.FetchExchangeRate(amount.currency, to)
	
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get exchange rate: %w", err)
	}

	amt, err := applyExchangeRate(amount, to, exchangeRate)
	if err != nil {
		return Amount{}, err
	}

	return amt, nil
}

// applyExchangeRate returns a new Amount representing the input multiplied by the ExchangeRate.
// The precision of the returned amount will match that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) (Amount, error) {
	decRate, err := ParseDecimal(fmt.Sprintf("%g", rate))
	if err != nil {
		return Amount{}, fmt.Errorf("could not convert exchange rate to decimal: %w", err)
	}
	converted := multiply(a.quantity, decRate)

	switch {
	case converted.precision > target.precision:
		converted.units /= pow10(converted.precision - target.precision)
	case converted.precision < target.precision:
		converted.units *= pow10(target.precision - converted.precision)
	}

	converted.precision = target.precision

	return Amount{
		currency: target,
		quantity: converted,
	}, nil
}

// multiply multiplies two Decimal values together to produce a new Decimal value.
func multiply(d1 Decimal, d2 Decimal) Decimal {
	dec := Decimal{
		units:     d1.units * d2.units,
		precision: d1.precision + d2.precision,
	}

	dec.simplify()
	return dec
}
