package main

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var (
	mainConfig  *Config
	mainConn    *Conn
	mainWG      sync.WaitGroup
	logger      *zap.Logger
	tickerDBMap TickerDBMap
)

func main() {

	var err error
	logger, err = SetupLogger()

	if err != nil {
		panic("couldn't start")
	}

	mainConfig, err = LoadDefaultConfig(configPath)
	logger.Info("loaded the main config.json")

	if err != nil {
		log.Fatal(err)
	}

	tickerDBMap, err = setupTickerDBs(mainConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("loaded the ticker databases")

	mainWG.Add(2)
	go SetupPoloniexConnection(&mainWG)
	logger.Info("setup connections to the exchanges")
	go startServer(mainConfig.ServerPort, &mainWG, tickerDBMap)
	logger.Info("started gRPC server", zap.Int("port", mainConfig.ServerPort))
	mainWG.Wait()
}
