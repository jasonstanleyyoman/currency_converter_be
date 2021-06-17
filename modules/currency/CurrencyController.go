package currency_converter

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/currency_be/utils"
)

type GinGonicCurrencyController struct {
	Service ICurrencyService
}

func (controller * GinGonicCurrencyController) GetAllRates(ctx * gin.Context) {
	symbol := ctx.DefaultQuery("symbol", "EUR")
	results, err :=controller.Service.GetAllRates(symbol)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenerateErrorWithMessage(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.GenerateOkResponse(
		GetAllRatesResponse{Rates: results,}))
}

func (controller * GinGonicCurrencyController) Convert(ctx * gin.Context) {
	sourceCurrency := ctx.Query("from")
	destCurrency := ctx.Query("to")

	if sourceCurrency == "" {
		ctx.JSON(http.StatusBadRequest, utils.GenerateErrorWithMessage("Missing from parameter"))
		return
	}

	if destCurrency == "" {
		ctx.JSON(http.StatusBadRequest, utils.GenerateErrorWithMessage("Missing to parameter"))
		return
	}

	amountInString := ctx.Query("amount")

	amountInFloat, errConvertingToFloat := strconv.ParseFloat(amountInString, 64)
	if errConvertingToFloat != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenerateErrorWithMessage("Amount is not a number"))
		return
	}

	result, errConverting := controller.Service.Convert(amountInFloat, sourceCurrency, destCurrency)

	if errConverting != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenerateErrorWithMessage(errConverting.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.GenerateOkResponse(
		ConvertResponse{
			Query: Query{
				From: sourceCurrency,
				To: destCurrency,
				Amount: amountInFloat,
			},
			Result: result,}))
}

func (controller * GinGonicCurrencyController) Update(ctx * gin.Context) {
	symbol := ctx.Query("symbol")
	amountInString := ctx.Query("amount")

	amountInFloat, errConvertingToFloat := strconv.ParseFloat(amountInString, 64)

	if errConvertingToFloat != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenerateErrorWithMessage("Amount is not a number"))
		return
	}

	err := controller.Service.UpdateRate(symbol, amountInFloat)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenerateErrorWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}

func NewGinGonicCurrencyController() GinGonicCurrencyController {
	currencyRepo := &CurrencyRepo{}
	currencyRepo.InitRepo()
	currencyService := &CurrencyService{Repo: currencyRepo}
	currencyController := GinGonicCurrencyController{Service: currencyService}

	return currencyController
}
