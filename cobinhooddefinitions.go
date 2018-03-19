package main

type cobinhoodWallet struct {
	Success bool `json:"success"`
	Result  struct {
		Balances []struct {
			Currency string `json:"currency"`
			Type     string `json:"type"`
			Total    string `json:"total"`
			OnOrder  string `json:"on_order"`
			Locked   bool   `json:"locked"`
			UsdValue string `json:"usd_value"`
			BtcValue string `json:"btc_value"`
		} `json:"balances"`
	} `json:"result"`
}

type cobinhoodOrders struct {
	Success bool `json:"success"`
	Result  struct {
		Limit     int `json:"limit"`
		Page      int `json:"page"`
		TotalPage int `json:"total_page"`
		Orders    []struct {
			ID            string      `json:"id"`
			TradingPairID string      `json:"trading_pair_id"`
			Side          string      `json:"side"`
			Type          string      `json:"type"`
			Price         string      `json:"price"`
			Size          string      `json:"size"`
			Filled        string      `json:"filled"`
			State         string      `json:"state"`
			Timestamp     int64       `json:"timestamp"`
			EqPrice       string      `json:"eq_price"`
			CompletedAt   interface{} `json:"completed_at"`
			TradingPair   string      `json:"trading_pair"`
		} `json:"orders"`
	} `json:"result"`
}

type cobinhoodOrder struct {
	Success bool `json:"success"`
	Result  struct {
		Order struct {
			ID          string `json:"id"`
			TradingPair string `json:"trading_pair"`
			State       string `json:"state"`
			Side        string `json:"side"`
			Type        string `json:"type"`
			Price       string `json:"price"`
			Size        string `json:"size"`
			Filled      string `json:"filled"`
			Timestamp   int64  `json:"timestamp"`
		} `json:"order"`
	} `json:"result"`
}

type cobinhoodTicker struct {
	Success bool `json:"success"`
	Result  struct {
		Ticker struct {
			TradingPairID  string `json:"trading_pair_id"`
			Timestamp      int64  `json:"timestamp"`
			Two4HHigh      string `json:"24h_high"`
			Two4HLow       string `json:"24h_low"`
			Two4HOpen      string `json:"24h_open"`
			Two4HVolume    string `json:"24h_volume"`
			LastTradePrice string `json:"last_trade_price"`
			HighestBid     string `json:"highest_bid"`
			LowestAsk      string `json:"lowest_ask"`
		} `json:"ticker"`
	} `json:"result"`
}

type cobinhoodPlaceOrder struct {
	TradingPairID string `json:"trading_pair_id"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	Price         string `json:"price"`
	Size          string `json:"size"`
}

type cobinhoodPlaceOrderResult struct {
	Success bool `json:"success"`
	Result  struct {
		Order struct {
			ID          string `json:"id"`
			TradingPair string `json:"trading_pair"`
			State       string `json:"state"`
			Side        string `json:"side"`
			Type        string `json:"type"`
			Price       string `json:"price"`
			Size        string `json:"size"`
			Filled      string `json:"filled"`
			Timestamp   int64  `json:"timestamp"`
			EqPrice     string `json:"eq_price"`
		} `json:"order"`
	} `json:"result"`
}
