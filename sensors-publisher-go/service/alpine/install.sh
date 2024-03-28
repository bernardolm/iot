#!/bin/bash

set -e

echo "$ workdir is $(pwd)"

# shellcheck source=/dev/null
source base.env

export ENV_PATH="/etc/$RC_SVCNAME"
export ERR_LOG="/var/log/$RC_SVCNAME/stderr.log"
export INSTALL_PATH="/usr/share/$RC_SVCNAME"
export LOG_PATH="/var/log/$RC_SVCNAME"
export OUTPUT_LOG="/var/log/$RC_SVCNAME/stdout.log"
export TMP_PATH="/tmp/$RC_SVCNAME"

function stop() {
    rc-service "$RC_SVCNAME" stop || true
    rc-update delete "$RC_SVCNAME" || true
    killall -9 "$RC_SVCNAME" 2>/dev/null || true
}

function clear() {
    rm -rf "${LOG_PATH:?}/*" "${INSTALL_PATH:?}/*" || \
        mkdir -p "$LOG_PATH" "$INSTALL_PATH"

    echo $ dest paths empty
}

function check_user() {
    id "$1" &>/dev/null
}

function add_user() {
    adduser -D "$USER_WORKER"
    addgroup "$USER_WORKER" root
}

function install() {
    echo $ checking user

    if ! check_user "$USER_WORKER"; then
        add_user
        echo "user $USER_WORKER added"
    fi

    echo $ moving files

    mv -f "$TMP_PATH/$RC_SVCNAME" "$INSTALL_PATH/$RC_SVCNAME"
    mv -f "$TMP_PATH/prod.env" "$ENV_PATH/config.env"
    mv -f "$TMP_PATH/autostart" "/etc/init.d/$RC_SVCNAME"

    echo $ fixing log files

    if [ ! -f "$ERR_LOG" ]; then
        touch "$ERR_LOG"
    fi

    if [ ! -f "$OUTPUT_LOG" ]; then
        touch "$OUTPUT_LOG"
    fi

    echo $ fixing permissions

    chmod a+rw "$ENV_PATH/config.env"
    chmod a+rw "$ERR_LOG"
    chmod a+rw "$OUTPUT_LOG"

    echo $ installation completed successfully
}

function check_install() {
    echo $ "files in ENV_PATH: $(ls -h "$ENV_PATH")"
    echo $ "files in INSTALL_PATH: $(ls -h "$INSTALL_PATH")"
    echo $ "files in LOG_PATH: $(ls -h "$LOG_PATH")"
    echo $ "files in TMP_PATH: $(ls -h "$TMP_PATH")"
}

function start() {
    rc-update add "$RC_SVCNAME" default
    rc-service "$RC_SVCNAME" start
    rc-service "$RC_SVCNAME" status
}

stop
clear
install
check_install
start

echo $ now it will only show recent logs. you can exit with ctrl+c.
tail -n30 -f "$ERR_LOG"
