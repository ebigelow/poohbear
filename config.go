package main

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

var (
	ErrLoadConfigErr = errors.New("couldn't load the config")
)

var (
	homeDir        = os.Getenv("HOME")
	poohbearDir    = path.Join(homeDir, ".poohbear")
	configPath     = path.Join(poohbearDir, "config.json")
	defaultDataDir = path.Join(poohbearDir, "data")
	defBTCTickerDb = path.Join(defaultDataDir, "bitfinex_ticker_btc.db")
	defLTCTickerDb = path.Join(defaultDataDir, "poloniex_ticker_ltc.db")
	defETHTickerDb = path.Join(defaultDataDir, "poloniex_ticker_eth.db")
	defLTCOrdersDb = path.Join(defaultDataDir, "poloniex_orders_ltc.db")
	defETHOrdersDb = path.Join(defaultDataDir, "poloniex_orders_eth.db")
)

type Config struct {
	ServerPort int           `json:"server_port"`
	DBPaths    ExchangeDBMap `json:"db_paths"`
}

func NewDefaultConfig() *Config {
	dbPaths := DefaultDBPathMap()
	return &Config{
		ServerPort: 12345,
		DBPaths:    dbPaths,
	}
}

func LoadDefaultConfig(path string) (*Config, error) {
	if _, err := os.Stat(path); err == nil {
		config, err := ParseConfig(path)
		return config, err
	} else if os.IsNotExist(err) {
		newConfig := NewDefaultConfig()
		file, createErr := os.Create(configPath)
		defer file.Close()
		if createErr == nil {
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			encoder.Encode(newConfig)
		}
		return newConfig, nil
	}
	return &Config{}, ErrLoadConfigErr
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

func DefaultDBPathMap() ExchangeDBMap {
	dbPaths := make(ExchangeDBMap)
	dbPaths["bitfinex"] = make(map[string]string)
	dbPaths["bitfinex"]["ticker_btc"] = defBTCTickerDb
	dbPaths["poloniex"] = make(map[string]string)
	dbPaths["poloniex"]["ticker_ltc"] = defLTCTickerDb
	dbPaths["poloniex"]["ticker_eth"] = defETHTickerDb
	dbPaths["poloniex"]["orders_ltc"] = defLTCOrdersDb
	dbPaths["poloniex"]["orders_eth"] = defETHOrdersDb
	return dbPaths
}
