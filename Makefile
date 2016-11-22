NAME      = ghl
GOVERSION = $(shell go version)
GOOS      = $(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH    = $(word 2,$(subst /, ,$(lastword $(GOVERSION))))
VERSION   = $(patsubst "%", %, $(lastword $(shell grep 'const Version' $(NAME).go)))
BUILD     = $(shell git rev-parse --verify HEAD)

.PHONY: build build-linux-amd64 build-linux-386 build-darwin-amd64 all installdeps test release clean

build: pkg/$(NAME)-$(GOOS)-$(GOARCH)

pkg/$(NAME)-$(GOOS)-$(GOARCH):
	go build -ldflags "-X main.build=$(BUILD)" -o pkg/$(NAME)-$(GOOS)-$(GOARCH) cmd/$(NAME)/$(NAME).go

build-linux-amd64:
	@$(MAKE) build GOOS=linux GOARCH=amd64

build-linux-386:
	@$(MAKE) build GOOS=linux GOARCH=386

build-darwin-amd64:
	@$(MAKE) build GOOS=darwin GOARCH=amd64

all: clean build-linux-amd64 build-linux-386 build-darwin-amd64

test:
	@go test -v $(shell glide nv)

installdeps:
	@glide install

release: test all
	@cd pkg && find . -type f | xargs -I{} sh -c "tar czvf {}.tar.gz {} && rm -f {}"
	@ghr $(VERSION) pkg/

clean:
	-rm -rf pkg/
