package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/clownpriest/poohbear/lib/poohbear"
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

func (db *TickerDB) AddTradeBlock(x poohbear.TradeBlock) error {
	data, err := x.Marshal()
	if err != nil {
		return err
	}
	db.store.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.mainBucket))
		err := b.Put([]byte(x.Timestamp), data)
		return err
	})
	return nil
}

func (db *TickerDB) GetTradeBlock(x []byte) *poohbear.TradeBlock {
	result := new(poohbear.TradeBlock)
	fmt.Println("inside GetTradeBlock")
	db.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.mainBucket)
		v := b.Get(x)
		result.Unmarshal(v)
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
	return result
}

func (db *TickerDB) GetTradeRange(start, end []byte) *poohbear.TradeBlockRange {
	startTime := []byte(start)
	endTime := []byte(end)
	result := new(poohbear.TradeBlockRange)
	db.store.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(db.mainBucket).Cursor()
		for k, v := c.Seek(startTime); k != nil && bytes.Compare(k, endTime) <= 0; k, v = c.Next() {
			var block poohbear.TradeBlock
			block.Unmarshal(v)
			result.Add(&block)
		}
		return nil
	})
	return result
}
