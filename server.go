package main

import (
	"fmt"
	"net"
	"sync"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func startServer(port int, wg *sync.WaitGroup, tickerMap TickerDBMap) {
	defer wg.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterTradeRangeServer(grpcServer, newTradeServer(tickerMap))
	grpcServer.Serve(lis)
}

type tradeServer struct {
	tickerMap TickerDBMap
}

func newTradeServer(tickerMap TickerDBMap) *tradeServer {
	s := new(tradeServer)
	s.tickerMap = tickerMap
	return s
}

func (ps *tradeServer) GetTradeRange(ctx context.Context, r *DateRange) (*TradeBlockRange, error) {
	var err error
	var result *TradeBlockRange
	switch r.Exchange {
	case "bitfinex":
		result, err = ps.handleBitfinexTicker(ctx, r)
	case "poloniex":
		result, err = ps.handlePoloniexTicker(ctx, r)
	}

	return result, err
}

func (ps *tradeServer) handlePoloniexTicker(ctx context.Context, r *DateRange) (*TradeBlockRange, error) {
	var err error
	var result *TradeBlockRange
	switch r.Pair {
	case "BTC_LTC":
		result, err = ps.tickerMap.DB("poloniex", "BTC_LTC").GetTradeRange([]byte(r.Start), []byte(r.End))
	case "BTC_ETH":
		result, err = ps.tickerMap.DB("poloniex", "BTC_ETH").GetTradeRange([]byte(r.Start), []byte(r.End))
	}
	return result, err
}

func (ps *tradeServer) handleBitfinexTicker(ctx context.Context, r *DateRange) (*TradeBlockRange, error) {
	var err error
	var result *TradeBlockRange
	switch r.Pair {
	case "BTC_USD":
		result, err = ps.tickerMap.DB("bitfinex", "BTC_USD").GetTradeRange([]byte(r.Start), []byte(r.End))
	}
	return result, err
}
