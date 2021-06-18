package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/currency_be/middleware"
	"github.com/jasonstanleyyoman/currency_be/routes"
	"github.com/joho/godotenv"
)

func main()  {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(0)
	}

	r := gin.Default()
	r.Use(middleware.Cors()...)
	routes.GinGonicRouter().InitRouter(r)
	r.Run()
}