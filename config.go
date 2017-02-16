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
	defLTCTickerDb = path.Join(defaultDataDir, "ticker_ltc.db")
	defETHTickerDb = path.Join(defaultDataDir, "ticker_eth.db")
	defLTCOrdersDb = path.Join(defaultDataDir, "orders_ltc.db")
	defETHOrdersDb = path.Join(defaultDataDir, "orders_eth.db")
)

type Config struct {
	ServerPort int               `json:"server_port"`
	DBPaths    map[string]string `json:"db_paths"`
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

func DefaultDBPathMap() map[string]string {
	db_paths := make(map[string]string)
	db_paths["ticker_ltc"] = defLTCTickerDb
	db_paths["ticker_eth"] = defETHTickerDb
	db_paths["orders_ltc"] = defLTCOrdersDb
	db_paths["orders_eth"] = defETHOrdersDb
	return db_paths
}
