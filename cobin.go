package main

import (
	"fmt"

	"github.com/cobinhoodGo"
)

const mycobinhoodAuth = `GET YOUR API KEY FROM YOUR COBINHOOD ACCOUNT`

func main() {
	//First set your API Key
	ch := new(cobinhoodgo.Cobin)
	ch.SetAPIKey(mycobinhoodAuth)

	//How to get your wallet
	wallet, err := cobinhoodgo.GetWallet(*ch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wallet)

	//How to get Ticker data
	tickersIwant := [2]string{"COB-BTC", "TRX-BTC"}
	tickers, err := cobinhoodgo.GetTicker(*ch, tickersIwant[0:2])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tickers)

	//How to get all Open Orders
	myOrders, err := cobinhoodgo.GetOpenOrders(*ch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myOrders)

	//How to check status of an Order
	for i := range myOrders {
		myOrderStatus, err := cobinhoodgo.GetOrderStatus(*ch, myOrders[i].ID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(myOrderStatus.ID + " " + myOrderStatus.State)
	}

	//How to place an Order
	var newOrder cobinhoodgo.PlaceOrderData
	newOrder.TradingPairID = "COB-BTC"
	newOrder.Side = "ask"
	newOrder.Type = "limit"
	newOrder.Price = 1
	newOrder.Size = 435

	myPlacedOrder, err := cobinhoodgo.PlaceOrder(*ch, newOrder)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myPlacedOrder)

}
