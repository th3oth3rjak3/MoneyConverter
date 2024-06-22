package money

// exchangeRates is a provider for currency exchange rate information.
type exchangeRates interface {
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}
