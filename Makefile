.PHONY: deps build doc fmt lint run run-fe test vendor_clean vendor_get vendor_update vet


GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test



deps:
	go get github.com/gin-gonic/gin
	go get golang.org/x/crypto/bcrypt
	go get gopkg.in/mgo.v2
	go get github.com/appleboy/gin-jwt
deps-fe:
	cd web && npm i
run:
	PORT=8080 go run ./main.go
run-fe:
	cd web && npm start