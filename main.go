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
	//gin.SetMode(gin.ReleaseMode)

	// Create a Gin router with the specified base path
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// Define routes under the "/binance" group
		binance := v1.Group("/binance")
		{
			binance.GET("/ticker/24hr", handlers.Binance().Get24HourTickerData)
			binance.GET("/ticker/24hr/:pair", handlers.Binance().GetTickerForPair)
			binance.GET("/ticker/24hr/gainers", handlers.Binance().Get24HourGainersTickerData)
			binance.GET("/ticker/24hr/gainers/pairs", handlers.Binance().Get24HourGainersPairs)

		} // Define routes under the "/binance" group
		bybit := v1.Group("/bybit")
		{
			bybit.GET("/ticker/24hr", handlers.Bybit().Get24HourTickerData)
			bybit.GET("/ticker/24hr/:pair", handlers.Bybit().GetTickerForPair)
			bybit.GET("/ticker/24hr/gainers", handlers.Bybit().Get24HourGainersTickerData)
			bybit.GET("/ticker/24hr/gainers/pairs", handlers.Bybit().Get24HourGainersPairs)

		}
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("127.0.0.1:8999")
}
