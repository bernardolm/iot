ifneq (,$(wildcard ./dev.env))
    include dev.env
    export
	ENV_FILE_PARAM = --env-file dev.env
endif

reset:
	@reset

air: reset
	@command -v air || go install github.com/cosmtrek/air@latest
	@command -v expect_unbuffer || sudo apt install expect
	@command -v goimports || go install golang.org/x/tools/cmd/goimports@latest
	@expect_unbuffer air -build.exclude_dir=tmp -c=.air.toml -d

start:
	@go run cmd/console/main.go

clear:
	@rm -rf bin/*
	@rm -rf dist/*

build: clear
	@go mod tidy
	@# GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/sensors-publisher-go-amd64 cmd/console/main.go
	@GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/sensors-publisher-go cmd/console/main.go
	@sleep 1
	@# command -v upx || sudo apt install upx-ucl
	@# upx --lzma -o dist/sensors-publisher-go bin/sensors-publisher-go

install: build
	@cp -f prod.env service/base.env service/${PLATFORM}/* bin/* dist/
	@ls -h dist
	@rsync -r ./dist/* "${SYSTEM_USER}@${SYSTEM_HOST}:/tmp/sensors-publisher-go"
	@ssh -t ${SYSTEM_USER}@${SYSTEM_HOST} "cd /tmp/sensors-publisher-go; sudo ./install.sh"

install-debian:
	@PLATFORM=debian make install

install-alpine:
	@PLATFORM=alpine make install

debug:
	ssh ${SYSTEM_USER}@${SYSTEM_HOST} "sudo -S tail -n100 -f /var/log/sensors-publisher-go/stderr.log"
