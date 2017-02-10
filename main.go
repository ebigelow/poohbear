package main

import (
	"log"
	"os"
	"sync"
)

var (
	mainConfig *Config
	mainConn   *Conn
	mainWG     sync.WaitGroup
	ltcTicker  *TickerDB
)

func main() {
	configPath := os.Args[1]
	var err error
	mainConfig, err = ParseConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	ltcTicker = LoadTickerDB("BTC_LTC", mainConfig.TickerDBLTC)
	defer ltcTicker.Close()

	mainWG.Add(1)
	go SetupConnection(&mainWG)
	mainWG.Wait()
}
