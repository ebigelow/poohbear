package main

import (
	"log"
	"os"
	"sync"
)

var (
	mainConfig  *Config
	mainConn    *Conn
	mainWG      sync.WaitGroup
	tickerDBMap map[string]*TickerDB
)

func main() {
	configPath := os.Args[1]
	var err error
	mainConfig, err = ParseConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	tickerDBMap = setupTickerDBs(mainConfig)

	mainWG.Add(2)
	go SetupConnection(&mainWG)
	go startServer(mainConfig.ServerPort, &mainWG, tickerDBMap)
	mainWG.Wait()

}
