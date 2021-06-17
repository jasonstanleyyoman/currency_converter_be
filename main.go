package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/currency_be/routes"
)

func main()  {
	r := gin.Default()
	routes.GinGonicRouter().InitRouter(r)
	r.Run()
}