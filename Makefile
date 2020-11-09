GOPATH := $(shell go env GOPATH | tr '\\' '/')
GOEXE := $(shell go env GOEXE)
GOHOSTOS := $(shell go env GOHOSTOS)
GOHOSTARCH := $(shell go env GOHOSTARCH)
GORELEASER := $(GOPATH)/bin/goreleaser
SOURCE_FILES := *.go public/* assets_vfsdata.go

all: build

$(GORELEASER):
	wget -qO- https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | BINDIR=$(GOPATH)/bin sh

build: $(SOURCE_FILES)
	go build

release-snapshot: $(GORELEASER) $(SOURCE_FILES)
	$(GORELEASER) release --snapshot --skip-publish --rm-dist

release: $(GORELEASER) $(SOURCE_FILES)
	$(GORELEASER) release --rm-dist

assets_vfsdata.go: assets_generate.go public/*
	go run assets_generate.go

clean:
	rm -rf dist tmp* assets_vfsdata.go

.PHONY: all build release release-snapshot clean
