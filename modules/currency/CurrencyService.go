package currency_converter

import (
	"errors"
)

type CurrencyService struct {
	Repo ICurrencyRepo
}

func (service * CurrencyService) Convert(amount float64, srcCurrencySymbol, destCurrencySymbol string) (float64 ,error) {
	srcBase, errSrc := service.Repo.GetRate(srcCurrencySymbol)
	destBase, errDst := service.Repo.GetRate(destCurrencySymbol)

	if errSrc != nil {
		return 0.0, errors.New("Unknown source currency symbol")
	}

	if errDst != nil {
		return 0.0, errors.New("Unknown destination currency symbol")
	}

	return (destBase / srcBase) * amount , nil
}

func (service * CurrencyService) GetAllRates(symbol string) (map[string]float64, error) {

	allCurrencySymbol := service.Repo.GetAllCurrencySymbol()
	rateBase, err := service.Repo.GetRate(symbol)
	eurBase, _ := service.Repo.GetRate("EUR")
	if err != nil {
		return make(map[string]float64), err
	}

	result := make(map[string]float64)

	for _, symbol := range allCurrencySymbol {
		rate, _ := service.Repo.GetRate(symbol)
		result[symbol] = rate * eurBase / (rateBase)
	}

	return result, nil
}