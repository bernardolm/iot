.EXPORT_ALL_VARIABLES:

SSH_PUB_KEY=$(shell cat ${HOME}/.ssh/id_ed25519.pub)
docker_container_name=alpine-sensors
docker_image_name="${GITHUB_USER}/${docker_container_name}:latest"

reset:
	@reset

air: reset
	@command -v air || go install github.com/cosmtrek/air@latest
	@command -v expect_unbuffer || sudo apt install -y expect
	@command -v goimports || go install golang.org/x/tools/cmd/goimports@latest
	@expect_unbuffer air -build.exclude_dir=tmp -c=.air.toml -d

start:
	@go run cmd/console/main.go

install:
	@./installer/start.sh

install-alpine:
	@ARCH=arm64 OS=alpine make install

install-debian:
	@ARCH=amd64 OS=debian make install

install-docker:
	@(docker stop ${docker_container_name} 2>/dev/null && sleep 1) || true
	@docker build --no-cache \
		--progress=tty \
		--build-arg SSH_PUB_KEY="${SSH_PUB_KEY}" \
		-f "docker/Dockerfile" \
		-t ${docker_image_name} \
		docker
	@docker run -d --rm --privileged \
		--name ${docker_container_name} \
		${docker_image_name} /bin/bash
	@ARCH=amd64 OS=alpine-docker make install

debug:
	@source dist/install.dev
	@ssh ${SYSTEM_USER}@${SYSTEM_HOST} "sudo -S tail -n50 -f ${LOG_FILEPATH}"
