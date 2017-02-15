.PHONY: protos

all: protos
	go install ./exchange/poloniex
	go install ./lib/poohbear
	go build

protos:
	protoc -I./protos --gofast_out=plugins=grpc:./lib/poohbear ./protos/trade.proto
	python -m grpc.tools.protoc -I./protos --python_out=./lib/python/poohbear --grpc_python_out=./lib/python/poohbear ./protos/trade.proto
