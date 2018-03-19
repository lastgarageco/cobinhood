package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const mycobinhoodAuth = `GET YOUR API KEY FROM COBINHOOD ACCOUNT`

func requestCobinhood(postType string, apiURL string, body io.Reader, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(postType, apiURL, body)
	req.Header.Add("Authorization", mycobinhoodAuth)
	req.Header.Add("nonce", strconv.FormatInt(time.Now().UnixNano()/1000000, 10))

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	myWallet := make(map[string]float64)
	myOrders := make(map[string]string)
	var myAskPrice float64
	var myBidPrice float64

	myWalletBalances := cobinhoodWallet{}
	requestCobinhood("GET", "https://api.cobinhood.com/v1/wallet/balances", nil, &myWalletBalances)

	for i := range myWalletBalances.Result.Balances {
		total, _ := strconv.ParseFloat(myWalletBalances.Result.Balances[i].Total, 64)
		onOrder, _ := strconv.ParseFloat(myWalletBalances.Result.Balances[i].OnOrder, 64)
		myWallet[myWalletBalances.Result.Balances[i].Currency] = total - onOrder
	}

	cobinhoodTickerCOBBTC := cobinhoodTicker{}
	requestCobinhood("GET", "https://api.cobinhood.com/v1/market/tickers/COB-BTC", nil, &cobinhoodTickerCOBBTC)

	lowestAsk, _ := strconv.ParseFloat(cobinhoodTickerCOBBTC.Result.Ticker.LowestAsk, 64)
	highestBid, _ := strconv.ParseFloat(cobinhoodTickerCOBBTC.Result.Ticker.HighestBid, 64)

	if (lowestAsk - highestBid) >= 0.00000025 {
		myAskPrice = lowestAsk - 0.00000001
		myBidPrice = myAskPrice - 0.00000020
	} else {
		myAskPrice = lowestAsk + 0.00000020
		myBidPrice = myAskPrice - 0.00000020
	}

	askOrder := &cobinhoodPlaceOrder{
		TradingPairID: "COB-BTC",
		Side:          "ask",
		Type:          "limit",
		Price:         strconv.FormatFloat(myAskPrice, 'f', 8, 64),
		Size:          strconv.FormatFloat(myWallet["COB"], 'f', 4, 64),
	}

	bidOrder := &cobinhoodPlaceOrder{
		TradingPairID: "COB-BTC",
		Side:          "bid",
		Type:          "limit",
		Price:         strconv.FormatFloat(myBidPrice, 'f', 8, 64),
		Size:          strconv.FormatFloat(myWallet["COB"], 'f', 4, 64),
	}

	orderJSON, _ := json.Marshal(askOrder)

	placedOrder := cobinhoodPlaceOrderResult{}
	requestCobinhood("POST", "https://api.cobinhood.com/v1/trading/orders", bytes.NewReader(orderJSON), &placedOrder)

	openOrders := cobinhoodOrders{}
	requestCobinhood("GET", "https://api.cobinhood.com/v1/trading/orders", nil, &openOrders)

	for i := range openOrders.Result.Orders {
		myOrders[openOrders.Result.Orders[i].ID] = openOrders.Result.Orders[i].Side
	}

	keepChecking := true
	for keepChecking {
		for id, side := range myOrders {
			fmt.Println("cheking status of order: " + id + "-" + side)
			orderStatus := cobinhoodOrder{}
			orderstatusURL := "https://api.cobinhood.com/v1/trading/orders/" + id

			requestCobinhood("GET", orderstatusURL, nil, &orderStatus)

			switch orderStatus.Result.Order.State {
			case "filled":
				if side == "ask" {
					orderJSON, _ = json.Marshal(bidOrder)
					placedOrder = cobinhoodPlaceOrderResult{}
					requestCobinhood("POST", "https://api.cobinhood.com/v1/trading/orders", bytes.NewReader(orderJSON), &placedOrder)
					myOrders[placedOrder.Result.Order.ID] = placedOrder.Result.Order.Side
				}
				delete(myOrders, id)
			case "cancelled":
				delete(myOrders, id)
			}
		}
		duration := time.Second * 5
		time.Sleep(duration)
		if len(myOrders) == 0 {
			keepChecking = false
		}
	}

}
