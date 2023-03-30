.PHONY: all
all: start

.PHONY: start
start:
	go run cmd/main.go

.PHONY: setup
setup:
	@go get -u github.com/gorilla/mux
	@go get -u github.com/teris-io/shortid
