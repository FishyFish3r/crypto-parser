package main

import (
	"context"
	"fmt"
	"strings"

	// grpc...
	selserv "github.com/FishyFish3r/crypto-parser/Parser/pkg/selserv"
	"github.com/PuerkitoBio/goquery"
	"google.golang.org/grpc"
)

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

type Coin struct {
	Name     string
	PriceKZT string
	PriceRUB string
}

func parseBinancePrice(url string, c selserv.SeleniumServerClient) string {
	resp, _ := c.GetHtml(context.Background(), &selserv.HtmlArgs{Url: url})

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(resp.GetHtml()))

	price := doc.Find(".showPrice").Text()

	return price
}

func parseCryptoruPrice(url string, c selserv.SeleniumServerClient) string {
	resp, _ := c.GetHtml(context.Background(), &selserv.HtmlArgs{Url: url})

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(resp.GetHtml()))

	price := doc.Find(".coin__price-value").Text()
	price = strings.TrimSpace(price)

	return price
}

func main() {
	conn, err := grpc.Dial(":61337", grpc.WithInsecure())

	if !checkError(err) {
		return
	}

	client := selserv.NewSeleniumServerClient(conn)

	for {
		var coin Coin

		coin.Name = "USDT"

		coin.PriceRUB = parseBinancePrice("https://www.binance.com/en/trade/USDT_RUB", client)
		coin.PriceKZT = parseCryptoruPrice("https://crypto.ru/usdt-kzt/", client)

		fmt.Printf("1 %s = \n\tRUB:%s\n\tKZT:%s\n", coin.Name, coin.PriceRUB, coin.PriceKZT)
	}

}
