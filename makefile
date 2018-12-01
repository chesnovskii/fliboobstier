GOCMD       = go
GOBUILD     = $(GOCMD) build
GOCLEAN     = $(GOCMD) clean
GOTEST      = $(GOCMD) test
GOGET       = $(GOCMD) get
GOINSTALL   = $(GOCMD) install
GOLINT      = $(GOPATH)/bin/golint

BINARY_NAME=bin

all: lint test build

build:
    $(GOBUILD) -o $(BINARY_NAME) -v

test:
    FLIBOOBSTIER_DEBUG=false FLIBOOBSTIER_TG_TOKEN='myLittleTestToken' $(GOTEST) -v ./...

clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)

run:
    $(GOBUILD) -o $(BINARY_NAME) -v ./...
    ./$(BINARY_NAME)

install:
    $(GOINSTALL) -v ./...

deps:
    $(GOGET) "github.com/Sirupsen/logrus"
    $(GOGET) "github.com/caarlos0/env"
    $(GOGET) "gopkg.in/yaml.v2"
    $(GOGET) "gopkg.in/telegram-bot-api.v4"

lint:
    $(GOGET) "github.com/golang/lint/golint"
    $(GOLINT) -min_confidence 0 -set_exit_status
