package main

import (
	"sync"

	"github.com/clownpriest/poohbear/exchange/poloniex"
	"github.com/k0kubun/pp"
)

var (
	mutex = sync.Mutex{}
)

func LTCTickerHandler(p []interface{}, n map[string]interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	var block poloniex.TradeBlock

	for _, value := range p {
		parsed := value.(map[string]interface{})
		if parsed["type"] == "newTrade" {
			t := poloniex.ParseTrade(parsed, "BTC_LTC")
			block.Timestamp = t.Timestamp
			block.Trades = append(block.Trades, &t)
			pp.Println(block)
			ltcTicker.AddTradeBlock(block)
		}
	}

}
