CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
COMMIT=`git rev-parse --short HEAD`
APP=fusion
REPO?=union-project/$(APP)
TAG?=latest
DEPS=$(shell go list ./... | grep -v /vendor/)

all: build

build:
	@cd cmd/$(APP) && go build -v -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT)" .

build-static:
	@cd cmd/$(APP) && go build -v -a -tags "netgo static_build" -installsuffix netgo -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT)" .

image: build-static
	@mkdir -p build
	@cp -r cmd/$(APP)/$(APP) build/
	@docker build -t $(REPO):$(TAG) .

release: image
	@docker push $(REPO):$(TAG)

check:
	@go vet -v $(DEPS)
	@golint $(DEPS)


test: build
	@go test -v ./...

clean:
	@rm -rf cmd/$(APP)/$(APP)
	@rm -rf build

.PHONY: build build-static image release test clean
