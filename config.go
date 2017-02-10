package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	TickerDBLTC string `json:"ticker_db_ltc"`
	OrdersDBLTC string `json:"orders_db_ltc"`
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
