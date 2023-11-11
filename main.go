package main

import (
	"github.com/cploutarchou/CryptoGainerAPI-Client/docs"
	"github.com/cploutarchou/CryptoGainerAPI-Client/handler"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func main() {
	key := os.Getenv("BINANCE_KEY")
	secret := os.Getenv("BINANCE_SECRET")
	cnf := parser.Config{Binance: &parser.Binance{
		ApiKey:    key,
		SecretKey: secret,
	}}
	docs.SwaggerInfo.BasePath = "/api/v1"
	parser_, _ := parser.New(cnf)
	handlers := handler.New(parser_)

	// Create a Gin router with the specified base path
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// Define routes under the "/binance" group
		binance := v1.Group("/binance")
		{
			binance.GET("/ticker/24hr", handlers.Binance().Get24HourTickerData)
			binance.GET("/ticker/24hr/:pair", handlers.Binance().GetTickerForPair)
		}
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
