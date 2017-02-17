import poohbear

client = poohbear.Client(port=12345)

result = client.up_to_now(exchange="poloniex", pair="BTC_LTC", minutes_ago=10)

print result.trades
