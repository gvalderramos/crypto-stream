.PHONY: all clean binance_watcher coingecko_watcher crypto_stream event_mng kracken_watcher
.DEFAULT_GOAL: all

BINARY_FOLDER=bin
CMD_FOLDER=cmd

all: binance_watcher coingecko_watcher crypto_stream event_mng kracken_watcher

clean:
	rm -rf ./bin

_go_mod_tidy:
	@go mod tidy

binance_watcher: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/binance_watcher $(CMD_FOLDER)/binance_watcher/*.go

coingecko_watcher: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/coingecko_watcher $(CMD_FOLDER)/coingecko_watcher/*.go

crypto_stream: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/crypto_stream $(CMD_FOLDER)/crypto_stream/*.go

event_mng: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/event_mng $(CMD_FOLDER)/event_mng/*.go

kracken_watcher: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/kracken_watcher $(CMD_FOLDER)/kracken_watcher/*.go

