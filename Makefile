.PHONY: all clean binance_watcher congecko_watcher crypto_stream db_mng event_mng kracken_watcher
.DEFAULT_GOAL: all

BINARY_FOLDER=bin
CMD_FOLDER=cmd

all: binance_watcher congecko_watcher crypto_stream db_mng event_mng kracken_watcher

clean:
	rm -rf ./bin

_go_mod_tidy:
	@go mod tidy

binance_watcher: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/binance_watcher $(CMD_FOLDER)/binance_watcher/main.go

congecko_watcher: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/congecko_watcher $(CMD_FOLDER)/congecko_watcher/main.go

crypto_stream: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/crypto_stream $(CMD_FOLDER)/crypto_stream/main.go

db_mng: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/db_mng $(CMD_FOLDER)/db_mng/main.go

event_mng: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/event_mng $(CMD_FOLDER)/event_mng/main.go

kracken_watcher: _go_mod_tidy
	@go build -o $(BINARY_FOLDER)/kracken_watcher $(CMD_FOLDER)/kracken_watcher/main.go

