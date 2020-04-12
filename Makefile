# Configure Go
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOSETTINGS=GO111MODULE=on GOPATH=$(GOPATH) GOBIN=$(GOBIN)
GOCMD=$(GOSETTINGS) go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) mod download

# Configure project
CMDDIR=$(GOBASE)/cmd
OUTPUTDIR=$(GOBIN)

# Define targets
all: clean test build

build: build-random

.PHONY: build-random
RANDOM_CMD=$(CMDDIR)/random/main.go
RANDOM_OUTPUT=$(OUTPUTDIR)/uci-impl-random
build-random:
	@echo "  >  Building Random solver..."
	$(GOBUILD) -i -o $(RANDOM_OUTPUT) $(RANDOM_CMD)

.PHONY: clean
clean:
	@echo "  >  Cleaning project..."
	$(GOCLEAN)
	rm -f $(RANDOM_OUTPUT)

.PHONY: test
test:
	@echo "  >  Running tests..."
	$(GOTEST) -v ./...

.PHONY: install
install:
	@echo "  >  Installing dependencies..."
	$(GOINSTALL)
