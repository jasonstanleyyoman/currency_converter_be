package currency_converter

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinGonicCurrencyController struct {
	Service ICurrencyService
}

func (controller * GinGonicCurrencyController) GetAllRates(ctx * gin.Context) {
	base := ctx.DefaultQuery("base", "EUR")
	results, err :=controller.Service.GetAllRates(base)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Currency not recognized")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": results,
	})
}

func (controller * GinGonicCurrencyController) Convert(ctx * gin.Context) {
	sourceCurrency := ctx.Query("from")
	destCurrency := ctx.Query("to")

	amountInString := ctx.Query("amount")

	amountInFloat, errConvertingToFloat := strconv.ParseFloat(amountInString, 64)
	if errConvertingToFloat != nil {
		ctx.String(http.StatusBadRequest, "Amount is not a number")
		return
	}

	result, errConverting := controller.Service.Convert(amountInFloat, sourceCurrency, destCurrency)

	if errConverting != nil {
		ctx.String(http.StatusBadRequest, errConverting.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (controller * GinGonicCurrencyController) Update(ctx * gin.Context) {

}

func NewGinGonicCurrencyController() GinGonicCurrencyController {
	currencyRepo := &CurrencyRepo{}
	currencyRepo.InitRepo()
	currencyService := &CurrencyService{Repo: currencyRepo}
	currencyController := GinGonicCurrencyController{Service: currencyService}

	return currencyController
}
