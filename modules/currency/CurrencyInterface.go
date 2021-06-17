package currency_converter

type ICurrencyRepo interface {
	GetRate(src string) (float64, error)
	UpdateRate(currencySymbol string, newBase float64) error
	GetAllCurrencySymbol() ([]string)
}

type ICurrencyService interface {
	Convert(amount float64, srcCurrencySymbol, destCurrencySymbol string) (float64 ,error)
	UpdateRate(currencySymbol string, newBase float64) (error)
	GetAllRates(symbol string) ([]Rate, error)
}

type Rate struct {
	Symbol string `json:"symbol"`
	Rate float64 `json:"rate"`
}
