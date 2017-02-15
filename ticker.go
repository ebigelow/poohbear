package main

func setupTickerDBs(config *Config) map[string]*TickerDB {
	ltcTicker := LoadTickerDB("BTC_LTC", mainConfig.TickerDBLTC)
	ethTicker := LoadTickerDB("BTC_ETH", mainConfig.TickerDBETH)

	tickerDBMap := make(map[string]*TickerDB)
	tickerDBMap["BTC_LTC"] = ltcTicker
	tickerDBMap["BTC_ETH"] = ethTicker
	return tickerDBMap
}
