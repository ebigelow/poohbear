import grpc
import trade_pb2


class PBClient(object):
    def __init__(self, port):
        self.channel = grpc.insecure_channel('localhost:{}'.format(port))
        self.trade_range_stub = trade_pb2.TradeRangeStub(self.channel)

    def get_range(self, trade_range):
        return self.trade_range_stub.GetTradeRange(trade_range)


client = PBClient(port=12345)

trade_range = trade_pb2.DateRange(pair="BTC_LTC",
                                  start="2017-02-14T17:38:59Z",
                                  end="2017-02-14T17:40:32Z")


result = client.get_range(trade_range=trade_range)

print result.trades
