BIN := /opt/hyperhash/bin/hh-core
PKG := ./cmd/core

.PHONY: build restart logs fmt lint test

build:
go build -trimpath -buildvcs=false -o $(BIN) $(PKG)

restart:
sudo systemctl restart hh-core

logs:
sudo journalctl -u hh-core -n 50 --no-pager

fmt:
go fmt ./...

lint:
golangci-lint run

test:
go test ./...
