package currency_converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type CurrencyRepo struct {
	CurrencyMapRates map[string]float64
}

func (repo * CurrencyRepo) GetRate(src string) (float64, error) {
	result, ok := repo.CurrencyMapRates[src]

	if !ok {
		return 0.0, errors.New("Unknown currency symol")
	}
	return result, nil
}

func (repo * CurrencyRepo) UpdateRate(currencySymbol string, newBase float64) {
	repo.CurrencyMapRates[currencySymbol] = newBase
}

func (repo * CurrencyRepo) GetAllCurrencySymbol() ([]string) {
	symbols := make([]string, 0)
	for key := range repo.CurrencyMapRates {
		symbols = append(symbols, key)
	}
	return symbols
}

type RateObject struct {
	Symbol string `json:"symbol"`
	Rate float64 `json:"rate"`
}

func (repo * CurrencyRepo) InitRepo() {
	repo.CurrencyMapRates = make(map[string]float64)

	byteValue, _ := ioutil.ReadFile("data/currency.json")

	allRates := make([]RateObject, 0)
	json.Unmarshal(byteValue, &allRates)

	for _, base := range allRates {
		repo.CurrencyMapRates[base.Symbol] = base.Rate
	}
	fmt.Println(repo.CurrencyMapRates)

}