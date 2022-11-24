RPI1_HOST=bernardo@192.168.1.11

start:
	go run cmd/console/main.go

build:
	# GOOS=linux GOARCH=amd64 go build -gccgoflags "-s -w" -ldflags "-s -w" -o console_amd64 cmd/console/main.go
	# -gccgoflags "-s -w"
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o console_arm64 cmd/console/main.go
	upx --best --lzma console_arm64
	scp console_arm64 ${RPI1_HOST}:/tmp
	ssh ${RPI1_HOST} 'env; /tmp/console_arm64'
