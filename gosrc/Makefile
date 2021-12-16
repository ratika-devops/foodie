#VERSION := $(shell git describe --always --long)
VERSION := v1.0.0

all: build

build-setup:
	@mkdir -p webui-docroot

web-build:
	@cd ../webui && make build

deps:
	@[ -f go.mod ] || { go mod init; }
	@go mod vendor
	@go mod tidy

swag:
	@echo "[ > ] Checking if swag is installed"
	@echo "Gopath is $$GOPATH"
	@set -x; ls -al $$GOPATH
	@[ `$$GOPATH/bin/swag -v | grep  version| wc -l | awk '{print $1}'` -eq 1 ] && { echo "[ PASS ] Golang swag is installed"; } || { echo "[ FAILED ] Golang swag is not installed"; echo "[ TRY ] Attempting to install Golang swag tool"; set -x; cd $$GOPATH;go get -v -u github.com/swaggo/swag/cmd/swag;  set +x;sync; [ `$$GOPATH/bin/swag -v | grep  version | wc -l | awk '{print $1}'` -eq 1 ] && { echo "[ PASS ] Golang swag is installed"; }; }
	@echo "[ > ] Updating swagger docs"
	@set -x; ls -al $$GOPATH
	@$$GOPATH/bin/swag init

build: web-build build-setup swag
	@set -x; go get github.com/markbates/pkger/cmd/pkger
	@set -x; rm -fR dist; mkdir -p dist
	@set -x; rm -fR webui-docroot; cp -a ../webui/dist webui-docroot
	@set -x; rm -f pkged.go server/pkged.go
	@set -x; pkger  -include /webui-docroot
	@set -x; sed 's/package main/package server/g' pkged.go > server/pkged.go
	@set -x; rm -f pkged.go
	@set -x; go mod vendor
	@set -x; go build -v -ldflags "-X main.version=${VERSION}" -o dist/foodies
	@set -x; ls -al dist/foodies

win-build: web-build build-setup swag
	@set -x; go get github.com/markbates/pkger/cmd/pkger
	@set -x; rm -fR dist; mkdir -p dist
	@set -x; rm -fR webui-docroot; cp -a ../webui/dist webui-docroot
	@set -x; rm -f pkged.go server/pkged.go
	@set -x; pkger  -include /webui-docroot
	@set -x; sed 's/package main/package server/g' pkged.go > server/pkged.go
	@set -x; rm -f pkged.go
	@set -x; go mod vendor
	@set -x; GOOS=windows GOARCH=amd64 go build -v -ldflags "-X main.version=${VERSION}" -o dist/foodies.exe
	@set -x; ls -al dist/foodies.exe

run: build 
	@chmod 755 dist/foodies
	@dist/foodies server

clean:
	@go clean -x

