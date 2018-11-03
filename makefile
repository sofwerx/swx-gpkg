# Go parameters
GOCMD=go
GOBIN=$(GOBIN)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=swxtorgi

# run these subcommands using make
all: test build

# builds go cmd line binary and outputs it to ~/go/bin/[filename]
# build location must be present in user path
build:
	$(GOBUILD) -o $(GOPATH)/bin/$(BINARY_NAME) -v

# run tests
test:
	$(GOTEST) -v ./...

# remove cmd line binaries
clean:
	$(GOCLEAN)
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

# go get
install:
	$(GOGET) github.com/edwardfward/swx-gpkg