package main

import (
	"sync"

	"github.com/clownpriest/poohbear/exchange/poloniex"
	"github.com/clownpriest/poohbear/lib/poohbear"
	"github.com/k0kubun/pp"
)

var (
	mutex = sync.Mutex{}
)

func LTCTickerHandler(p []interface{}, n map[string]interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	block := poohbear.TradeBlock{Pair: "BTC_LTC"}
	for _, value := range p {
		parsed := value.(map[string]interface{})
		if parsed["type"] == "newTrade" {
			t := poloniex.ParseTrade(parsed, "BTC_LTC")
			block.Timestamp = t.Timestamp
			block.Trades = append(block.Trades, &t)
			pp.Println(block)
			tickerDBMap["BTC_LTC"].AddTradeBlock(block)
		}
	}
}

func ETHTickerHandler(p []interface{}, n map[string]interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	block := poohbear.TradeBlock{Pair: "BTC_ETH"}
	for _, value := range p {
		parsed := value.(map[string]interface{})
		if parsed["type"] == "newTrade" {
			t := poloniex.ParseTrade(parsed, "BTC_ETH")
			block.Timestamp = t.Timestamp
			block.Trades = append(block.Trades, &t)
			pp.Println(block)
			tickerDBMap["BTC_ETH"].AddTradeBlock(block)
		}
	}
}
