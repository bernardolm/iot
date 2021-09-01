
install-alpine:
	apk update && apk upgrade

	apk fix --no-cache --reinstall \
		git \
		git-bash-completion \
		git-perl \
		git-zsh-completion \
		libffi \
		pcre2 \
		perl-error \
		perl-git

	apk add --no-cache \
		alpine-sdk \
		binutils \
		build-base \
		cargo \
		cython \
		file \
		fortify-headers \
		g++ \
		gcc \
		gcc6 \
		libc-dev \
		libffi-dev \
		libxml2-dev \
		libxslt-dev \
		libxslt-dev\
		linux-headers \
		make \
		musl-dev \
		openssl-dev \
		python2-dev \
		python3-dev

	curl -s https://bootstrap.pypa.io/get-pip.py | python3 -

	pip uninstall -y cffi pycparser
	pip install --upgrade cffi pycparser bcrypt
	pip install -r requirements.txt

install:
	@echo 'installing app...'
	# @curl -sL https://github.com/bernardolm/iot/archive/refs/heads/master.zip -o master.zip
	# @unzip -qq -o master.zip
	# @cp -r iot-master/* ../
	# @rm -rf iot-master master.zip
	@echo 'finish app install'

setup:
	@pip install .

kill:
	@echo 'killing zombie process...'
	-@for l in $(ps aux | grep src/ds_sensors_mqtt/main.py | grep -v grep | awk '{print $2}'); do \
		echo "killing process $l"; \
		kill -9 $l; \
	done

start: setup
	@echo 'starting app...'
	@python3 src/ds_sensors_mqtt/main.py

supervisor:
	@supervisord -c supervisord.conf

supervisor-program: install setup kill start

mosquitto:
	@docker run --rm -it -p 1883:1883 -p 9001:9001 \
		-v $(PWD)/mosquitto.conf:/mosquitto/config/mosquitto.conf \
		--name mosquitto eclipse-mosquitto
