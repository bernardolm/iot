
install-rpi-deps:
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

	apk add \
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
