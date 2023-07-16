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
	@ssh ${RPI_USER}@${RPI_HOST} "mkdir -p /tmp/sensors-publisher-go"
	@ssh -t ${RPI_USER}@${RPI_HOST} "eval ${SUDO} killall -9 sensors-publisher-go || true"
	@scp config.env \
		service/* \
		bin/sensors-publisher-go \
		${RPI_USER}@${RPI_HOST}:/tmp/sensors-publisher-go
	@ssh -t ${RPI_USER}@${RPI_HOST} "/tmp/sensors-publisher-go/install.sh"

debug:
	@ssh ${RPI_USER}@${RPI_HOST} "$(SUDO) journalctl -u sensors_publisher_go_service"
