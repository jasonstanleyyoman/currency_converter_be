package routes

import (
	"github.com/gin-gonic/gin"
	currency_converter "github.com/jasonstanleyyoman/currency_be/modules/currency"
)

func InitCurrencyRoute(g * gin.RouterGroup) {
	currencyController := currency_converter.NewGinGonicCurrencyController()

	currencyGroup := g.Group("/currency")
	{
		currencyGroup.GET("/", currencyController.GetAllRates)
		currencyGroup.GET("/convert", currencyController.Convert)
		currencyGroup.POST("/update", currencyController.Update)
	}
}
