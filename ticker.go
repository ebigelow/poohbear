package main

func setupTickerDBs(config *Config) (TickerDBMap, error) {
	tickerDBMap := make(TickerDBMap)

	bitfinexBTCTicker, err := LoadTickerDB("BTC_USD", mainConfig.DBPaths.For("bitfinex", "ticker_btc"))
	if err != nil {
		return tickerDBMap, err
	}

	poloniexLTCTicker, err := LoadTickerDB("BTC_LTC", mainConfig.DBPaths.For("poloniex", "ticker_ltc"))
	if err != nil {
		return tickerDBMap, err
	}

	poloniexETHTicker, err := LoadTickerDB("BTC_ETH", mainConfig.DBPaths.For("poloniex", "ticker_eth"))
	if err != nil {
		return tickerDBMap, err
	}

	tickerDBMap.Set("bitfinex", "BTC_USD", bitfinexBTCTicker)
	tickerDBMap.Set("poloniex", "BTC_LTC", poloniexLTCTicker)
	tickerDBMap.Set("poloniex", "BTC_ETH", poloniexETHTicker)

	return tickerDBMap, nil
}
