import grpc
import trade_pb2
import trade_time


class Client(object):
    def __init__(self, port):
        self.channel = grpc.insecure_channel('localhost:{}'.format(port))
        self.trade_range_stub = trade_pb2.TradeRangeStub(self.channel)

    def get_trade_range(self, trade_range):
        return self.trade_range_stub.GetTradeRange(trade_range)

    def up_to_now(self, pair="", days_ago=0,
                  hours_ago=0, minutes_ago=0,
                  seconds_ago=0):

        time_span = trade_time.up_to_now(pair=pair,
                                         days_ago=days_ago,
                                         hours_ago=hours_ago,
                                         minutes_ago=minutes_ago)
        result = self.trade_range_stub.GetTradeRange(time_span)
        return result
