ifneq (,$(wildcard ./config.env))
    include config.env
    export
endif

SUDO=

ifneq (${RPI_USER},"root")
	SUDO="sudo "
endif

start:
	@go run cmd/console/main.go

build:
	@GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/sensors-publisher-go cmd/console/main.go
	@upx --best --lzma bin/sensors-publisher-go

install: build
	@rm -rf dist/*
	@cp -f bin/* service/${PLATFORM}/* config.env dist/
	@ls -ah dist
	@scp dist/* ${RPI_USER}@${RPI_HOST}:/tmp/sensors-publisher-go
	@ssh -t ${RPI_USER}@${RPI_HOST} "/tmp/sensors-publisher-go/install.sh"

install-debian:
	@PLATFORM=debian make install

install-alpine:
	@PLATFORM=alpine make install

debug:
	@ssh ${RPI_USER}@${RPI_HOST} tail -f /var/log/sensors-publisher-go/stderr.log
