package main

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser"
	"os"
)

func main() {
	key := os.Getenv("BINANCE_KEY")
	secret := os.Getenv("BINANCE_SECRET")
	cnf := parser.Config{Binance: &parser.Binance{
		ApiKey:    key,
		SecretKey: secret,
	}}
	parser, _ := parser.New(cnf)
	tickers, _ := parser.Binance().Get24HourTickerData()
	data, _ := json.MarshalIndent(tickers, "", "    ")
	fmt.Println(string(data))

}
