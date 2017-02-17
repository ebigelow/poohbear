package main

type ExchangeDBMap map[string]map[string]string

func (exchMap ExchangeDBMap) For(exchange, feed string) string {
	return exchMap[exchange][feed]
}

type TickerDBMap map[string]map[string]*TickerDB

func (tickMap TickerDBMap) DB(exchange, pair string) *TickerDB {
	return tickMap[exchange][pair]
}

func (tickMap TickerDBMap) Set(exchange string, pair string, db *TickerDB) {
	if _, ok := tickMap[exchange]; ok {
		tickMap[exchange][pair] = db
	} else {
		tickMap[exchange] = make(map[string]*TickerDB)
		tickMap[exchange][pair] = db
	}

}
