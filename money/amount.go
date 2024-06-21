package money

// Amount represents a quantity of money in a specific currency.
type Amount struct {
	quantity Decimal
	currency Currency
}
