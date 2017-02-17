package main

import (
	"log"
	"sync"

	"github.com/jcelliott/turnpike"
)

type Conn struct {
	client *turnpike.Client
}

func SetupPoloniexConnection(wg *sync.WaitGroup) {
	defer wg.Done()
	c, err := turnpike.NewWebsocketClient(turnpike.JSON, PoloniexWSAddress, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.JoinRealm("realm1", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Subscribe("BTC_LTC", nil, PoloniexLTCTickerHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Subscribe("BTC_ETH", nil, PoloniexETHTickerHandler)
	if err != nil {
		log.Fatal(err)
	}

	mainConn := new(Conn)
	mainConn.client = c
	select {}

}
