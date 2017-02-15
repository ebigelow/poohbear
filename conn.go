package main

import (
	"log"
	"sync"

	"github.com/clownpriest/poohbear/exchange/poloniex"
	"github.com/jcelliott/turnpike"
)

type Conn struct {
	client *turnpike.Client
}

func SetupConnection(wg *sync.WaitGroup) {
	defer wg.Done()
	c, err := turnpike.NewWebsocketClient(turnpike.JSON, poloniex.WSAddress, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.JoinRealm("realm1", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Subscribe("BTC_LTC", nil, LTCTickerHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Subscribe("BTC_ETH", nil, ETHTickerHandler)
	if err != nil {
		log.Fatal(err)
	}

	mainConn := new(Conn)
	mainConn.client = c
	select {}

}
