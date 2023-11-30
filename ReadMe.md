[![CircleCI](https://dl.circleci.com/status-badge/img/circleci/8xWFaF83UGzqoLTaHXvpfe/4s2uuhZWD7k4mWZyLdaqAG/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/circleci/8xWFaF83UGzqoLTaHXvpfe/4s2uuhZWD7k4mWZyLdaqAG/tree/main)



# CryptoGainerAPI-Client Main Package Documentation

The `main` package in the CryptoGainerAPI-Client project is the entry point for the application. It initializes the necessary components, sets up API routes, and starts the web server to serve cryptocurrency data using the Gin framework.

## Dependencies

This package relies on the following external libraries and modules:

- `github.com/cploutarchou/CryptoGainerAPI-Client/docs`: Contains Swagger API documentation information.
- `github.com/cploutarchou/CryptoGainerAPI-Client/handler`: Defines API request handlers.
- `github.com/cploutarchou/CryptoGainerAPI-Client/parser`: Provides configuration parsing and API client setup.
- `github.com/gin-gonic/gin`: Gin framework for building web APIs.
- `github.com/swaggo/files` and `github.com/swaggo/gin-swagger`: Swagger documentation generator for APIs.
- `os`: Standard Go package for working with the operating system environment variables.

## Environment Variables

Before running the application, make sure to set the following environment variables:

- `BINANCE_KEY`: Binance API key for authentication.
- `BINANCE_SECRET`: Binance API secret key for authentication.

## Initialization

The `main` function performs the following initialization steps:

1. Reads the Binance API key and secret key from environment variables.
2. Initializes the application configuration using the `parser` package.
3. Creates API request handlers using the `handler` package.
4. Sets up the Gin router and defines API routes.

## API Routes

The application defines API routes under the `/api/v1` base path, grouped by cryptocurrency exchange:

### Binance API Routes

- `/api/v1/binance/ticker/24hr`: Get 24-hour ticker data for all pairs.
- `/api/v1/binance/ticker/24hr/:pair`: Get ticker data for a specific pair.
- `/api/v1/binance/ticker/24hr/gainers`: Get 24-hour gainers ticker data.
- `/api/v1/binance/ticker/24hr/gainers/pairs`: Get pairs with the highest 24-hour gains.

### Bybit API Routes

- `/api/v1/bybit/ticker/24hr`: Get 24-hour ticker data for all pairs.
- `/api/v1/bybit/ticker/24hr/:pair`: Get ticker data for a specific pair.
- `/api/v1/bybit/ticker/24hr/gainers`: Get 24-hour gainers ticker data.
- `/api/v1/bybit/ticker/24hr/gainers/pairs`: Get pairs with the highest 24-hour gains.

## Swagger Documentation

Swagger documentation for the API is available at `/docs/*any`. You can access the API documentation using a web browser or API client by visiting this route. It provides detailed information about the available endpoints and their usage.

## Running the Application

To start the CryptoGainerAPI-Client application, the web server listens on `127.0.0.1:8999`. You can access the API and documentation through this address.

```go
router.Run("127.0.0.1:8999")
