
install-rpi-deps:
	apk update && apk upgrade
	apk fix --no-cache --reinstall \
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
		git \
		git-bash-completion \
		git-perl \
		git-zsh-completion \
		libc-dev \
		libffi \
		libffi-dev \
		libxml2-dev \
		libxslt-dev \
		libxslt-dev\
		linux-headers \
		make \
		musl-dev \
		openssl-dev \
		pcre2 \
		perl-error \
		perl-git \
		python2-dev \
		python3-dev

	curl -s https://bootstrap.pypa.io/get-pip.py | python3 -

	export PATH='/usr/path/:${PATH}'

	pip uninstall -y cffi pycparser
	pip install --upgrade cffi pycparser bcrypt
