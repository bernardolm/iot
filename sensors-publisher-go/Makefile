ifneq (,$(wildcard ./config.env))
    include config.env
    export
endif

WORKING_DIR=/usr/local/sensor-publisher-go
BIN_FILE=console_arm

start:
	@go run cmd/console/main.go

build:
	@GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ${BIN_FILE} cmd/console/main.go
	@upx --best --lzma ${BIN_FILE}

install: build
	@ssh ${RPI1_USER}@${RPI1_HOST} 'rm -rf ${WORKING_DIR} || true; mkdir -p ${WORKING_DIR} || true'
	@scp config.env \
		sensor-publisher-go.service \
		service_install.sh \
		${BIN_FILE} \
		${RPI1_USER}@${RPI1_HOST}:${WORKING_DIR}
	@ssh ${RPI1_USER}@${RPI1_HOST} '${WORKING_DIR}/service_install.sh'

debug:
	@ssh ${RPI1_USER}@${RPI1_HOST} 'journalctl -u sensor_publisher_go_service'
