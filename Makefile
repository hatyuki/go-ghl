NAME        = ghl
GOVERSION   = $(shell go version)
GOOS        = $(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH      = $(word 2,$(subst /, ,$(lastword $(GOVERSION))))
VERSION     = $(patsubst "%", %, $(lastword $(shell grep 'const Version' $(NAME).go)))
BUILD       = $(shell git rev-parse --verify HEAD)
DEVTOOLS    = devtools

.PHONY: build all test installdeps release clean

build: pkg/$(NAME)-$(GOOS)-$(GOARCH)

pkg/$(NAME)-$(GOOS)-$(GOARCH):
	go build -ldflags "-X main.build=$(BUILD)" -o pkg/$(NAME)-$(GOOS)-$(GOARCH) cmd/$(NAME)/$(NAME).go

all: clean
	@$(MAKE) build GOOS=linux GOARCH=amd64
	@$(MAKE) build GOOS=linux GOARCH=386
	@$(MAKE) build GOOS=darwin GOARCH=amd64

test:
	@go test -v $(shell glide nv)

$(DEVTOOLS)/glide:
		@echo "Installing glide"
		@mkdir -p $(DEVTOOLS)
		@wget -O - "https://ghal.ga/masterminds/glide?os=$(GOOS)&arch=$(GOARCH)" | tar xvz
		@mv $(GOOS)-$(GOARCH)/glide $(DEVTOOLS)
		@rm -rf $(GOOS)-$(GOARCH)

installdeps:
	@which glide >/dev/null 2>&1 || $(MAKE) $(DEVTOOLS)/glide
	@PATH=$(DEVTOOLS):$(PATH) glide install

release: all
	@cd pkg && find . -type f | xargs -I{} sh -c "tar czvf {}.tar.gz {} && rm -f {}"
	@ghr --username hatyuki --token $(GITHUB_TOKEN) --replace $(VERSION) pkg/

clean:
	-rm -rf pkg/ devtools/ $(GOOS)-$(GOARCH)/
