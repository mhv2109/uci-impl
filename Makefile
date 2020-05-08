# Configure Go
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOSETTINGS=GO111MODULE=on GOPATH=$(GOPATH) GOBIN=$(GOBIN)
GOCMD=$(GOSETTINGS) go
GOBUILD=$(GOCMD) build -gcflags=-m
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -gcflags=-m
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) mod download
GODOC=$(GOSETTINGS) godoc

# Configure project
CMDDIR=$(GOBASE)/cmd
OUTPUTDIR=$(GOBIN)

# Define targets
all: clean test build

build: build-random build-minimax

.PHONY: build-random
RANDOM_CMD=$(CMDDIR)/random/main.go
RANDOM_OUTPUT=$(OUTPUTDIR)/mhv2109-uci-random
build-random:
	@echo "  >  Building Random solver..."
	$(GOBUILD) -i -o $(RANDOM_OUTPUT) $(RANDOM_CMD)

.PHONY: build-minimax
MINIMAX_CMD=$(CMDDIR)/minimax/main.go
MINIMAX_OUTPUT=$(OUTPUTDIR)/mhv2109-uci-minimax
build-minimax:
	@echo "  >  Building Minimax solver..."
	$(GOBUILD) -i -o $(MINIMAX_OUTPUT) $(MINIMAX_CMD)

.PHONY: clean
clean:
	@echo "  >  Cleaning project..."
	$(GOCLEAN)
	rm -f $(RANDOM_OUTPUT)

.PHONY: build-mocks
build-mocks:
	@echo " > Building mocks..."
	mockgen -source=internal/solver/solver.go -destination=internal/solver/mock/solver.go -package=mock
	mockgen -source=internal/handler/output.go -destination=internal/handler/mock/output.go -package=mock

.PHONY: test
test: build-mocks
	@echo "  >  Running tests..."
	$(GOTEST) -v -race ./...

.PHONY: benchmark-minimax
benchmark-minimax: build-mocks
	@echo "  >  Benchmarking Minimax algo..."
	$(GOTEST) -benchmem -cpuprofile=minimax_cpu.prof -memprofile=minimax_mem.prof github.com/mhv2109/uci-impl/internal/solver/minimax -bench="."

.PHONY: install
install:
	@echo "  >  Installing dependencies..."
	$(GOINSTALL)

.PHONY: doc
doc:
	@echo "  >  Running godoc..."
	$(GODOC)
