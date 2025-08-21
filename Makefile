APP=$(shell basename $(shell git remote get-url origin) | tr '[:upper:]' '[:lower:]')
REGISTRY=lotarion
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETARCH=amd64

format:
	gofmt -s -w ./

get_deps:
	go get

lint:
	golint

test:
	go test -v

build: format get_deps
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -ldflags "-X="github.com/Lotarion/GlobalLogic_DEVOPS101_kbot/cmd.appVersion=${VERSION} -o bin/kbot

image:
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

clean:
	rm -fr bin