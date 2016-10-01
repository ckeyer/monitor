PWD := $(shell pwd)
PKG := github.com/ckeyer/monitor
APP := monitor
DEV_IMAGE := ckeyer/dev

GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

GO := CGO_ENABLED=1 go

build-permetheus-image:
	cd docker-iamges && ./build.sh

dev:
	docker run --rm -it \
	 -v $(PWD):/opt/gopath/src/$(PKG) \
	 -w /opt/gopath/src/$(PKG) \
	 -p 5000:5000 \
	 $(DEV_IMAGE) bash

build:
	$(GO) build -o bundles/$(APP)

test-client:
	$(GO) test ./client