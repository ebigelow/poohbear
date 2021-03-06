package main

import (
	"sync"

	"github.com/k0kubun/pp"
)

var (
	mutex = sync.Mutex{}
)

func PoloniexLTCTickerHandler(p []interface{}, n map[string]interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	block := TradeBlock{Pair: "BTC_LTC"}
	for _, value := range p {
		parsed := value.(map[string]interface{})
		if parsed["type"] == "newTrade" {
			t := ParseTradePoloniex(parsed, "BTC_LTC")
			block.Timestamp = t.Timestamp
			block.Trades = append(block.Trades, &t)
			pp.Println(block)
			tickerDBMap.DB("poloniex", "BTC_LTC").AddTradeBlock(block)
		}
	}
}

func PoloniexETHTickerHandler(p []interface{}, n map[string]interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	block := TradeBlock{Pair: "BTC_ETH"}
	for _, value := range p {
		parsed := value.(map[string]interface{})
		if parsed["type"] == "newTrade" {
			t := ParseTradePoloniex(parsed, "BTC_ETH")
			block.Timestamp = t.Timestamp
			block.Trades = append(block.Trades, &t)
			pp.Println(block)
			tickerDBMap.DB("poloniex", "BTC_ETH").AddTradeBlock(block)
		}
	}
}
