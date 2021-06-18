package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/currency_converter_be/middleware"
	"github.com/jasonstanleyyoman/currency_converter_be/routes"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors()...)
	routes.GinGonicRouter().InitRouter(r)
	r.Run()
}
