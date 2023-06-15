package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

type Coin struct {
	Name      string
	PriceUSDT string
	PriceRUB  string
}

func main() {

	data := []byte("https://www.binance.com/en/trade/USDT_RUB")

	resp, err := http.Post("localhost:61137", "text/plain", bytes.NewBuffer(data))

	if !checkError(err) {
		return
	}

	data, _ = ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}
