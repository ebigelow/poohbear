package main

func setupTickerDBs(config *Config) (map[string]*TickerDB, error) {
	tickerDBMap := make(map[string]*TickerDB)

	ltcTicker, err := LoadTickerDB("BTC_LTC", mainConfig.DBPaths["ticker_ltc"])
	if err != nil {
		return tickerDBMap, err
	}

	ethTicker, err := LoadTickerDB("BTC_ETH", mainConfig.DBPaths["ticker_eth"])
	if err != nil {
		return tickerDBMap, err
	}

	tickerDBMap["BTC_LTC"] = ltcTicker
	tickerDBMap["BTC_ETH"] = ethTicker
	return tickerDBMap, nil
}
