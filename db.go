package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/clownpriest/poohbear/exchange/poloniex"
)

type TickerDB struct {
	store      *bolt.DB
	mainBucket []byte
}

func LoadTickerDB(pair, path string) *TickerDB {
	tickerDB := &TickerDB{store: new(bolt.DB)}
	err := tickerDB.Open(pair, path)
	if err != nil {
		return nil
	}
	return tickerDB
}

func (db *TickerDB) Open(pair, path string) error {
	tickerDB, openErr := bolt.Open(path, 0600, nil)

	if openErr != nil {
		log.Fatal(openErr)
		return openErr
	}

	db.store = tickerDB
	db.mainBucket = []byte(pair + "_ticker")
	err := db.store.Update(func(tx *bolt.Tx) error {
		_, updateErr := tx.CreateBucketIfNotExists(db.mainBucket)
		if updateErr != nil {
			log.Fatal(updateErr)
			return updateErr
		}
		return updateErr
	})

	return err
}

func (db *TickerDB) Close() {
	db.store.Close()
}

func (db *TickerDB) AddTradeBlock(x poloniex.TradeBlock) error {
	data, err := x.Marshal()
	if err != nil {
		return err
	}
	db.store.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.mainBucket))
		err := b.Put(x.Timestamp, data)
		return err
	})
	return nil
}

func (db *TickerDB) GetTradeBlock(x []byte) poloniex.TradeBlock {
	var result poloniex.TradeBlock
	db.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.mainBucket)
		v := b.Get(x)
		result.Unmarshal(v)
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
	return result
}
