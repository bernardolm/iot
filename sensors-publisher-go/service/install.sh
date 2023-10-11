#!/bin/sh

echo "starting service install"

CURRENT_DIR=$(dirname "$(readlink -f "$0")")

. "${CURRENT_DIR}/config.env"

if [ "$RPI_USER" != "root" ]; then
    SUDO="sudo "
fi

DISTRO=$(cat /etc/*-release | grep '^ID' | cut -d= -f2)

echo "distro is $DISTRO"

case "$DISTRO" in
    "debian" | "ubuntu")    . "${CURRENT_DIR}/debian/install.sh" ;;
    "alpine" )              . "${CURRENT_DIR}/alpine/install.sh" ;;
    *)                      echo "distro '$DISTRO' not found" ;;
esac

echo "finish service install"
