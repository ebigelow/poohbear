all:
	protoc -I./protos --gofast_out=plugins=grpc:./exchange/poloniex ./protos/poloniex.proto
	go install ./exchange/poloniex
	go build
