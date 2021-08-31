
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

setup:
	@curl -sL https://github.com/bernardolm/iot/archive/refs/heads/master.zip -o master.zip
	@unzip -qq -o master.zip
	@cp -r iot-master/* ../
	@rm -rf iot-master
	@pip install -r requirements.txt
	@pip install .

start: setup
	@python3 src/ds_sensors_mqtt/init.py

supervisor:
	@supervisord -c supervisord.conf
