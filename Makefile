GOPATH=$(CURDIR)
GO=~/Downloads/go/bin/go
GO_BINDATA=~/go/bin/go-bindata

test: 
	$(GO) test -cover -coverprofile=coverage.out

coverage:
	$(GO) tool cover -html=coverage.out

data:
	$(GO_BINDATA) -o server/data.go -pkg server data/...

deps:
	$(GO) get -u github.com/twinj/uuid
	$(GO) get -u github.com/gorilla/websocket

.PHONY: deps test coverage