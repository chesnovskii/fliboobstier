GOCMD       = go
GOBUILD     = $(GOCMD) build
GOCLEAN     = $(GOCMD) clean
GOTEST      = $(GOCMD) test
GOGET       = $(GOCMD) get
GOINSTALL   = $(GOCMD) install

BINARY_NAME=bin

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	FLIBOOBSTIER_DEBUG=false FLIBOOBSTIER_TG_TOKEN='myLittleTestToken' $(GOTEST) -count=50 ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

install:
	$(GOINSTALL) -v ./...

deps:
	$(GOGET) "github.com/sirupsen/logrus"
	$(GOGET) "github.com/caarlos0/env"
	$(GOGET) "gopkg.in/yaml.v2"
	$(GOGET) "gopkg.in/telegram-bot-api.v4"
	$(GOGET) "github.com/mattn/go-sqlite3"
	$(GOGET) "github.com/stretchr/testify"
	$(GOGET) "github.com/google/uuid"
