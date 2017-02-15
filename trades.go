package main

func (trade *TradeBlockRange) Add(block *TradeBlock) {
	trade.Trades = append(trade.Trades, block)
}
