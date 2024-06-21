package money

const (
	// ErrInvalidCurrencyCode is returned when the currency code is not a valid ISO 4217 code.
	ErrInvalidCurrencyCode = Error("invalid currency code")
)

// Currency represents a currency code that follows the ISO 4217 standard. e.g. USD, EUR, etc.
type Currency struct {
	code      string
	precision byte
}

// ParseCurrency parses a string representation of a currency code and returns a Currency.
func ParseCurrency(code string) (Currency, error) {
	if err := validateCurrencyCode(code); err != nil {
		return Currency{}, err
	}

	switch code {
	case "IRR": // Iranian Rial
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND": // Chinese Yuan, Vietnamese Dong
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND": // Bahraini Dinar, Iraqi Dinar, Kuwaiti Dinar, Libyan Dinar, Omani Rial, Tunisian Dinar
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}

// validateCurrencyCode checks if the currency code is a valid ISO 4217 code.
// It currently uses a naive approach and only checks the length and character range.
func validateCurrencyCode(code string) error {
	if len(code) != 3 {
		return ErrInvalidCurrencyCode
	}

	for _, c := range code {
		if c < 'A' || c > 'Z' {
			return ErrInvalidCurrencyCode
		}
	}

	return nil
}
