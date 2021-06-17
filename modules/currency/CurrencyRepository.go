package currency_converter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type CurrencyRepo struct {
	CurrencyMapRates map[string]float64
}

func (repo * CurrencyRepo) GetRate(src string) (float64, error) {
	result, ok := repo.CurrencyMapRates[src]

	if !ok {
		return 0.0, errors.New("Unknown currency symbol")
	}
	return result, nil
}

func (repo * CurrencyRepo) UpdateRate(currencySymbol string, newBase float64) (error) {
	if _, ok := repo.CurrencyMapRates[currencySymbol]; !ok {
		return errors.New("Unknown currency symbol")
	}

	repo.CurrencyMapRates[currencySymbol] = newBase
	return nil
	
}

func (repo * CurrencyRepo) GetAllCurrencySymbol() ([]string) {
	symbols := make([]string, 0)
	for key := range repo.CurrencyMapRates {
		symbols = append(symbols, key)
	}
	return symbols
}

func (repo * CurrencyRepo) InitRepo() {
	repo.CurrencyMapRates = make(map[string]float64)

	byteValue, _ := ioutil.ReadFile("data/currency.json")

	allRates := make([]Rate, 0)
	json.Unmarshal(byteValue, &allRates)

	for _, rate := range allRates {
		repo.CurrencyMapRates[rate.Symbol] = rate.Rate
	}
}