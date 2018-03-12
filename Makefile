.PHONY: deps build doc fmt lint run test vendor_clean vendor_get vendor_update vet


GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get



deps:
	$(GOGET) github.com/gin-gonic/gin
	$(GOGET) gopkg.in/mgo.v2
	$(GOGET) github.com/appleboy/gin-jwt;
run:
	$(GOCMD) run ./main.go