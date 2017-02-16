.PHONY: protos build

all: protos build

build:
	go build

protos:
	protoc -I./protos --gofast_out=plugins=grpc:. ./protos/trade.proto
	python -m grpc.tools.protoc -I./protos --python_out=./lib/python/poohbear --grpc_python_out=./lib/python/poohbear ./protos/trade.proto

deps:
	go get google.golang.org/grpc
	go get github.com/jcelliott/turnpike
	go get github.com/boltdb/bolt/...
	go get github.com/k0kubun/pp
