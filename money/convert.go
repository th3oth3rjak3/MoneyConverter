package money

// ExchangeRate represents a rate to convert from one currency to another.
type ExchangeRate Decimal

// Convert applies an exchange rate to convert an input amount to a target currency.
func Convert(amount Amount, to Currency) (Amount, error) {
	dec, err := ParseDecimal("2")
	if err != nil {
		return Amount{}, nil
	}

	exchangeRate := ExchangeRate(dec)

	amt := applyExchangeRate(amount, to, exchangeRate)
	return amt, nil
}

// applyExchangeRate returns a new Amount representing the input multiplied by the ExchangeRate.
// The precision of the returned amount will match that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {

	converted := multiply(a.quantity, Decimal(rate))

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
	}
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
