package main

import (
	"fmt"
	"time"

	"github.com/cobinhoodGo"
)

const mycobinhoodAuth = `GET AND ENTER YOUR API KEY FROM YOUR COBINHOOD ACCOUNT`

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
	tickers, err := cobinhoodgo.GetTicker(*ch, "COB-BTC", "TRX-BTC")
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
	newOrder.Side = "bid"
	newOrder.Type = "limit"
	newOrder.Price = 0.00000500
	newOrder.Size = 435

	myPlacedOrder, err := cobinhoodgo.PlaceOrder(*ch, newOrder)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myPlacedOrder)

	//How to cancel an Order
	time.Sleep(5 * time.Second) //this time is not necessary is just to make a small pause after the order we just palced.
	isCancel, err := cobinhoodgo.CancelOrder(*ch, myPlacedOrder.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isCancel)

}
