run: deps
	NEO4J_URL=http://localhost:7474/db/data GOPATH=$(CURDIR)/.go go run truckman.go

deps:
	GOPATH=$(CURDIR)/.go go get -d

build: deps
	GOPATH=$(CURDIR)/.go go build -o build/truckman truckman.go
