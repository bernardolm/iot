ifneq (,$(wildcard ./config.env))
    include config.env
    export
endif

start:
	@go run cmd/console/main.go

clear:
	@rm -rf bin/*
	@rm -rf dist/*

build: clear
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/sensors-publisher-go-amd64 cmd/console/main.go
	@GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/sensors-publisher-go cmd/console/main.go
	@sleep 1
	@upx --lzma -o dist/sensors-publisher-go bin/sensors-publisher-go

install: build
	@cp -f service/${PLATFORM}/* dist/
	@ls -ah dist
	@ssh -t ${SYSTEM_USER}@${SYSTEM_HOST} "mkdir -p /tmp/sensors-publisher-go"
	@scp dist/* ${SYSTEM_USER}@${SYSTEM_HOST}:/tmp/sensors-publisher-go
	@ssh -t ${SYSTEM_USER}@${SYSTEM_HOST} "/tmp/sensors-publisher-go/install.sh"

install-debian:
	@PLATFORM=debian make install

install-alpine:
	@PLATFORM=alpine make install

debug:
	@ssh ${SYSTEM_USER}@${SYSTEM_HOST} sudo -S tail -n1000 -f /var/log/sensors-publisher-go/stderr.log
