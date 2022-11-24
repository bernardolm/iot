ifneq (,$(wildcard ./config.env))
    include config.env
    export
endif

start:
	go run cmd/console/main.go

build:
	@GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o console_arm cmd/console/main.go
	@upx --best --lzma console_arm
	@ssh ${RPI1_USER}@${RPI1_HOST} 'env; rm -rf /tmp/console_* || true'
	@scp console_arm ${RPI1_USER}@${RPI1_HOST}:/tmp
	@scp Automation_Custom_Script.sh ${RPI1_USER}@${RPI1_HOST}:/boot
	@rm console_arm*
	@ssh root@${RPI1_HOST} '/boot/dietpi/dietpi-autostart 14'
