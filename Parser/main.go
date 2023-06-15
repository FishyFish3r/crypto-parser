package main

import (
	"context"
	"fmt"

	// grpc...
	selserv "github.com/FishyFish3r/crypto-parser/Parser/pkg/selserv"
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
	Name      string
	PriceUSDT string
	PriceRUB  string
}

func main() {
	conn, err := grpc.Dial(":61337", grpc.WithInsecure())

	if !checkError(err) {
		return
	}

	client := selserv.NewSeleniumServerClient(conn)

	client.GetHtml(context.Background(), &selserv.HtmlArgs{Url: "google.com"})
}
