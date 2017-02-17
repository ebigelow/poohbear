.PHONY: protos build deps

all: protos build

build:
	go build

install: deps
	go install
	@echo "creating ~/.poohbear directory"
	@if [ ! -d $(HOME)/.poohbear/data ]; then\
		mkdir -p $(HOME)/.poohbear/data;\
	fi
	@echo "done!"

clean:
	@echo "deleting ~/.poohbear directory"
	@if [ -d $(HOME)/.poohbear ]; then\
		rm -rf $(HOME)/.poohbear;\
	fi
	@echo "done!"

protos:
	protoc -I./protos --gofast_out=plugins=grpc:. ./protos/trade.proto
	python -m grpc.tools.protoc -I./protos --python_out=./lib/python/poohbear --grpc_python_out=./lib/python/poohbear ./protos/trade.proto

deps:
	go get google.golang.org/grpc
	go get github.com/jcelliott/turnpike
	go get github.com/boltdb/bolt/...
	go get github.com/k0kubun/pp
	go get go.uber.org/zap
	go get github.com/bitfinexcom/bitfinex-api-go
