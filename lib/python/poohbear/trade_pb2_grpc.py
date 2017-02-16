# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc
from grpc.framework.common import cardinality
from grpc.framework.interfaces.face import utilities as face_utilities

import trade_pb2 as trade__pb2


class TradeRangeStub(object):

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetTradeRange = channel.unary_unary(
        '/main.TradeRange/GetTradeRange',
        request_serializer=trade__pb2.DateRange.SerializeToString,
        response_deserializer=trade__pb2.TradeBlockRange.FromString,
        )


class TradeRangeServicer(object):

  def GetTradeRange(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TradeRangeServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetTradeRange': grpc.unary_unary_rpc_method_handler(
          servicer.GetTradeRange,
          request_deserializer=trade__pb2.DateRange.FromString,
          response_serializer=trade__pb2.TradeBlockRange.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'main.TradeRange', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))