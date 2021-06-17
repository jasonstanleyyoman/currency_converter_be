package currency_converter

import (
	"errors"
	"fmt"
)

type CurrencyService struct {
	Repo ICurrencyRepo
}

func (service * CurrencyService) Convert(amount float64, srcCurrencySymbol, destCurrencySymbol string) (ConversionResult ,error) {
	srcBase, errSrc := service.Repo.GetRate(srcCurrencySymbol)
	destBase, errDst := service.Repo.GetRate(destCurrencySymbol)

	if errSrc != nil {
		return ConversionResult{}, errors.New(fmt.Sprintf("Unknown source currency symbol : %s", srcCurrencySymbol))
	}

	if errDst != nil {
		return ConversionResult{}, errors.New(fmt.Sprintf("Unknown destination currency symbol : %s", destCurrencySymbol))
	}

	srcLong, _ := service.Repo.GetLongSymbol(srcCurrencySymbol)
	dstLong, _ := service.Repo.GetLongSymbol(destCurrencySymbol)

	return ConversionResult{
		From: srcCurrencySymbol,
		FromLongSymbol: srcLong,
		To: destCurrencySymbol,
		ToLongSymbol: dstLong,
		Amount: amount,
		Result: (destBase / srcBase) * amount,
	}, nil
}

func (service * CurrencyService) UpdateRate(currencySymbol string, newBase float64) (error) {
	return service.Repo.UpdateRate(currencySymbol, newBase)
}

func (service * CurrencyService) ConvertMultiple(amount float64, 
srcCurrencySymbol string, destCurrencySymbols []string) ([]ConversionResult, error) {
	results := make([]ConversionResult, 0)
	for _, destCurrencySymbol := range destCurrencySymbols {
		result, err := service.Convert(amount, srcCurrencySymbol, destCurrencySymbol)
		if err != nil {
			return results, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (service * CurrencyService) GetAllRates(symbol string) ([]Rate, error) {

	allCurrencySymbol := service.Repo.GetAllCurrencySymbol()
	rateBase, err := service.Repo.GetRate(symbol)
	eurBase, _ := service.Repo.GetRate("EUR")
	if err != nil {
		return make([]Rate, 0), err
	}

	result := make([]Rate, 0)

	for _, symbol := range allCurrencySymbol {
		rate, _ := service.Repo.GetRate(symbol)
		longSymbol, _ := service.Repo.GetLongSymbol(symbol)
		result = append(result, Rate{
			Symbol: symbol,
			Rate: rate * eurBase / (rateBase),
			LongSymbol: longSymbol,
		})
	}

	return result, nil
}