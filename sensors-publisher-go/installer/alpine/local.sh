#!/bin/bash

# set -e

echo "$ sh: installing..."

TMP_PATH=$(pwd)

echo "$ sh: workdir is ${TMP_PATH}"
echo "$ sh: USER_PWD=${USER_PWD}"

# shellcheck source=/dev/null
source .env

INSTALL_PATH="/usr/share/$RC_SVCNAME"
LOG_PATH="/var/log/$RC_SVCNAME.log"

function stop() {
    sudo rc-service "$RC_SVCNAME" stop || true
    sudo rc-update delete "$RC_SVCNAME" || true
    killall -9 "$RC_SVCNAME" 2>/dev/null || true
}

function clear() {
    echo "$ sh: cleaning..."

    sudo rm -rf "$INSTALL_PATH" "$LOG_PATH"
    sudo mkdir -m a+r "$INSTALL_PATH"

    check_install
}

function check_user() {
    id "$1" &>/dev/null
}

function manage_user() {
    sudo deluser --remove-home "$WORKER_USER" || true
    sudo adduser -g "$WORKER_USER" -s /bin/ash -G root -S -D "$WORKER_USER"
    echo "$ sh: user $WORKER_USER added"
    cat < /etc/passwd | grep "$WORKER_USER"
}

function install() {
    echo "$ sh: user checking..."

    $CREATE_USER_SKIP || manage_user

    echo "$ sh: moving files..."

    sudo cp --remove-destination "$TMP_PATH/.env" "$INSTALL_PATH/.env"
    sudo cp --remove-destination "$TMP_PATH/$RC_SVCNAME" "$INSTALL_PATH/$RC_SVCNAME"
    sudo cp --remove-destination "$TMP_PATH/autostart" "/etc/init.d/$RC_SVCNAME"

    echo "$ sh: giving permissions..."

    sudo chmod a+rx "/etc/init.d/$RC_SVCNAME"
    sudo chmod -R 777 "$INSTALL_PATH"
    $CREATE_USER_SKIP sudo chown -R "$WORKER_USER" "$INSTALL_PATH"

    sudo rc-update add "$RC_SVCNAME" default

    check_install

    echo "$ sh: installation completed"
}

function check_install() {
    echo "$ sh: files in INSTALL_PATH ($INSTALL_PATH): [$(ls -aCmN $INSTALL_PATH)]"
    echo "$ sh: files in TMP_PATH ($TMP_PATH): [$(ls -aCmN $TMP_PATH)]"
}

function start() {
    sudo rc-service "$RC_SVCNAME" restart || true
    sudo rc-service "$RC_SVCNAME" status
}

stop
clear
install
start

echo "$ sh: now it will only show recent logs. you can exit with ctrl+c."
tail -n30 -f "$LOG_PATH"

/bin/bash
