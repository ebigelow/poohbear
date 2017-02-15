package poloniex

import (
	"log"
	"strconv"
	"time"

	"github.com/clownpriest/poohbear/lib/poohbear"
)

const timeLayout = "2006-01-02 15:04:05"

var (
	amount    float64
	rate      float64
	total     float64
	date      time.Time
	timestamp = make([]byte, 8)
)

func ParseTrade(x map[string]interface{}, pair string) poohbear.Trade {
	t := poohbear.Trade{}
	t.Pair = pair
	var err error
	data := x["data"].(map[string]interface{})
	amount, err = strconv.ParseFloat(data["amount"].(string), 32)
	if err != nil {
		log.Fatal(err)
	}
	t.Amount = float32(amount)

	rate, err = strconv.ParseFloat(data["rate"].(string), 32)
	if err != nil {
		log.Fatal(err)
	}
	t.Rate = float32(rate)

	total, err = strconv.ParseFloat(data["total"].(string), 32)
	if err != nil {
		log.Fatal(err)
	}
	t.Total = float32(total)

	date, err := time.Parse(timeLayout, data["date"].(string))
	if err != nil {
		log.Fatal(err)
	}

	t.Timestamp = date.Format(time.RFC3339)
	return t
}
