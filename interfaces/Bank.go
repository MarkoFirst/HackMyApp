package interfaces

type PrivatBankExchangeRates struct {
	Date            string
	Bank            string
	BaseCurrency    int
	BaseCurrencyLit string
	ExchangeRate    []struct {
		BaseCurrency   string
		Currency       string
		SaleRateNB     float64
		PurchaseRateNB float64
	}
}

type Currency struct {
	Usd Rates
	Euro Rates
	Rub Rates
}

type Rates struct {
	In  float64
	Out float64
}
