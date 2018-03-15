.PHONY: deps build doc fmt lint run run-fe test vendor_clean vendor_get vendor_update vet
export GOPATH=${PWD}


GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOSRC=${GOPATH}/src/golang-server

deps:
	go get ./...
deps-fe:
	cd web && npm i
run:
	PORT=8080 go run ${GOSRC}/main.go
run-fe:
	cd web && npm start