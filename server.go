package main

import (
	"fmt"
	"net"
	"sync"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func startServer(port int, wg *sync.WaitGroup, tickerMap map[string]*TickerDB) {
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
	tickerMap map[string]*TickerDB
}

func newTradeServer(tickerMap map[string]*TickerDB) *tradeServer {
	s := new(tradeServer)
	s.tickerMap = tickerMap
	return s
}

func (ps *tradeServer) GetTradeRange(ctx context.Context, r *DateRange) (*TradeBlockRange, error) {
	var result *TradeBlockRange
	switch r.Pair {
	case "BTC_LTC":
		result = ps.tickerMap[r.Pair].GetTradeRange([]byte(r.Start), []byte(r.End))
	case "BTC_ETH":
		result = ps.tickerMap[r.Pair].GetTradeRange([]byte(r.Start), []byte(r.End))
	}
	return result, nil
}
