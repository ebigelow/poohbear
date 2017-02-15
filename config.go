package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerPort  int    `json:"server_port"`
	TickerDBLTC string `json:"ticker_db_ltc"`
	OrdersDBLTC string `json:"orders_db_ltc"`
	TickerDBETH string `json:"ticker_db_eth"`
	OrdersDBETH string `json:"orders_db_eth"`
}

func ParseConfig(path string) (*Config, error) {
	configFile, openErr := os.Open(path)
	if openErr != nil {
		return &Config{}, openErr
	}
	decoder := json.NewDecoder(configFile)
	config := Config{}
	decoder.Decode(&config)
	return &config, nil
}
