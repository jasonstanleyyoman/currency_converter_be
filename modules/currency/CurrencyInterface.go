package currency_converter

type ICurrencyRepo interface {
	GetRate(src string) (float64, error)
	GetLongSymbol(src string) (string, error)
	UpdateRate(currencySymbol string, newBase float64) error
	GetAllCurrencySymbol() []string
}

type ICurrencyService interface {
	Convert(amount float64, srcCurrencySymbol,
		destCurrencySymbol string) (ConversionResult, error)
	ConvertMultiple(amount float64,
		srcCurrencySymbol string, destCurrencySymbols []string) ([]ConversionResult, error)
	UpdateRate(currencySymbol string, newBase float64) error
	GetAllRates(symbol string) ([]Rate, error)
}

type Rate struct {
	Symbol     string  `json:"symbol"`
	Rate       float64 `json:"rate"`
	LongSymbol string  `json:"long"`
}

type ConversionResult struct {
	From           string  `json:"from"`
	FromLongSymbol string  `json:"from_long"`
	To             string  `json:"to"`
	ToLongSymbol   string  `json:"to_long"`
	Amount         float64 `json:"amount"`
	Result         float64 `json:"result"`
}
